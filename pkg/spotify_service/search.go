package spotify_service

import (
	"context"
	"encoding/json"
	"fmt"

	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func (s *SpotifyRetrievalServer) GetSearchResult(ctx context.Context, req *spotifyv1.GetSearchResultRequest) (*spotifyv1.GetSearchResultResponse, error) {
	var token oauth2.Token
	json.Unmarshal(req.Token, &token)
	//fmt.Println("token recieved by spotify: ", token)
	client := s.GetClient(&token)
	var mask spotify.SearchType
	for _, searchtype := range req.Types {
		spotify_search_type, err := GetSearchType(searchtype)
		if err != nil {
			return nil, err
		}
		mask = mask | spotify_search_type
	}
	fmt.Print(req.Query)
	search_result, err := client.Search(req.Query, spotify.SearchTypeTrack|spotify.SearchTypeAlbum)
	if err != nil {
		return nil, err
	}
	encoded_result, err := json.Marshal(search_result)
	if err != nil {
		return nil, err
	}
	return &spotifyv1.GetSearchResultResponse{
		SearchResponse: encoded_result,
	}, nil
}
