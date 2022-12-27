package spotify_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	session_managementv1 "github.com/anyuan-chen/colormatch/gen/proto/go/session_management/v1"
	sharedv1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	"github.com/anyuan-chen/colormatch/pkg/color_service"
	"golang.org/x/oauth2"
)

type SpotifyApi struct {
	Session_Manager session_managementv1.SessionManagementServiceClient
	Spotify_Service spotifyv1.SpotifyImageColorMatchingServiceClient
}
type Validity struct {
	Validity bool   `json:"valid"`
	Reason   string `json:"reason,omitempty"`
}
func (s *SpotifyApi) PingTokenValidity(w http.ResponseWriter, r *http.Request) {
	token, err := GetToken(s, r)
	fmt.Println("token retrieved", token)
	if err != nil {
		resp, err := json.Marshal(Validity{Validity: false, Reason: "token not found in request"})
		if err != nil {
			http.Error(w, "bad encoding json for validity", http.StatusInternalServerError)
			return
		}
		w.Write(resp)
		return
	}
	token_json, err := json.Marshal(token)
	if err != nil {
		resp, _ := json.Marshal(Validity{Validity: false, Reason: "token not found"})
		w.Write(resp)
		return
	}

	fmt.Println("token retrieved json version", token)
	res, err := s.Spotify_Service.PingTokenValidity(context.Background(), &spotifyv1.PingTokenValidityRequest{
		Token: token_json,
	})
	fmt.Println("call complete without fatal errors")
	if err != nil || !res.Valid {
		fmt.Println("ping token to spotify auth error: ", err)
		resp, err := json.Marshal(Validity{Validity: false})
		if err != nil {
			http.Error(w, "bad encoding json for validity", http.StatusInternalServerError)
			return
		}
		w.Write(resp)
		return
	}
	resp, err := json.Marshal(Validity{Validity: true})
	if err != nil {
		http.Error(w, "bad encoding json for validity", http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func (s *SpotifyApi) GetColorSummary(w http.ResponseWriter, r *http.Request) {
	token, err := GetToken(s, r)
	if err != nil {
		http.Error(w, "bad credentials", http.StatusBadRequest)
		return
	}
	token_json, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "bad token", http.StatusBadRequest)
		return
	}
	query := r.FormValue("query")
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "bad parameters", http.StatusBadRequest)
		return
	}
	search_type := r.FormValue("search_type")
	background_color_strings := r.Form["background_colors"]
	highlight_color_strings := r.Form["highlight_colors"]
	background_palette, err := StringArrayToPalette(background_color_strings)
	if err != nil {
		http.Error(w, "bad background color hex code", http.StatusBadRequest)
		return
	}
	highlight_palette, err := StringArrayToPalette(highlight_color_strings)
	if err != nil {
		http.Error(w, "bad highlight color hex code", http.StatusBadRequest)
		return
	}
	color_metadata_request := &spotifyv1.GetColorMetadataForSpotifyAssetRequest{
		Token:            token_json,
		Query:            query,
		SearchType:       GetSearchType(search_type),
		BackgroundColors: &background_palette,
		HighlightColors:  &highlight_palette,
	}
	color_metadata, err := s.Spotify_Service.GetColorMetadataForSpotifyAsset(context.Background(), color_metadata_request)
	if err != nil {
		http.Error(w, err.Error()+"this was the fault of the color retrieval failing", http.StatusBadRequest)
		return
	}
	final_color_data, err := json.Marshal(color_metadata)
	if err != nil {
		http.Error(w, "unparsable final color data"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(final_color_data)
}

func StringArrayToPalette(strings []string) (sharedv1.Palette, error) {
	palette := make([]*sharedv1.Color, len(strings))
	for _, str := range strings {
		rgb, err := color_service.HexToRgb(str)
		if err != nil {
			return sharedv1.Palette{}, err
		}
		palette = append(palette, &sharedv1.Color{
			R: int32(rgb.R),
			G: int32(rgb.G),
			B: int32(rgb.B),
			A: int32(rgb.A),
		})
	}
	return sharedv1.Palette{
		Color: palette,
	}, nil
}
func GetSearchType(search_type string) spotifyv1.SearchType {
	type_map := make(map[string]spotifyv1.SearchType)
	type_map["album"] = spotifyv1.SearchType_SEARCH_TYPE_ALBUM
	type_map["artist"] = spotifyv1.SearchType_SEARCH_TYPE_ARTIST
	return type_map[search_type]
}

func GetToken(s *SpotifyApi, r *http.Request) (oauth2.Token, error) {
	cookie, err := r.Cookie("spotify_session_id")
	if err != nil {
		fmt.Println(err)
		return oauth2.Token{}, err
	}
	retrieval_request := &session_managementv1.RetrieveRequest{
		Ciphertext: cookie.Value,
	}
	fmt.Println("cookie value: ", cookie.Value)
	retrieval_response, err := s.Session_Manager.Retrieve(context.Background(), retrieval_request)
	fmt.Println(retrieval_response.Token)
	if err != nil {
		fmt.Println(err)
		return oauth2.Token{}, err
	}
	var token oauth2.Token
	json.Unmarshal(retrieval_response.Token, &token)
	fmt.Println("final returned token", token)
	return token, nil
}
