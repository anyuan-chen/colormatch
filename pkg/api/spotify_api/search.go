package spotify_api

import (
	"context"
	"encoding/json"
	"net/http"

	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
)

func (s *SpotifyApi) Search(w http.ResponseWriter, r *http.Request) {
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
	params := r.URL.Query()
	query := params.Get("query")
	form_search_types := params["type"]
	search_types := make([]spotifyv1.SearchType, 0, len(form_search_types))
	for _, search_type := range form_search_types {
		search_types = append(search_types, GetSearchType(search_type))
	}
	req := spotifyv1.GetSearchResultRequest{
		Token: token_json,
		Query: query,
		Types: search_types,
	}
	res, err := s.Spotify_Service.GetSearchResult(context.Background(), &req)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "unencodable data"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(res.SearchResponse)
}
