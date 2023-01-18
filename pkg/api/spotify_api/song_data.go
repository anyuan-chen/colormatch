package spotify_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
)

func (s *SpotifyApi) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	artistIds := params["artist_ids"]
	trackIds := params["track_ids"]
	genres := params["genres"]
	recommendationCount, err := strconv.Atoi(params.Get("recommendation_count"))
	if err != nil {
		http.Error(w, "invalid recommendation count", http.StatusUnauthorized)
		return
	}
	token, err := GetToken(s, r)
	if err != nil {
		http.Error(w, "no token", http.StatusUnauthorized)
		return
	}
	encoded_token, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	request := spotifyv1.GetRecommendationsRequest{
		Token:               encoded_token,
		ArtistIds:           artistIds,
		TrackIds:            trackIds,
		Genres:              genres,
		RecommendationCount: int32(recommendationCount),
	}
	resp, err := s.Spotify_Service.GetRecommendations(context.Background(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp.Recommendations)
	fmt.Println("successfully finished writing audio recommendations")
}

func (s *SpotifyApi) GetTrackAudioAnalysis(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	token, err := GetToken(s, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token_json, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req := spotifyv1.GetTrackAudioAnalysisRequest{
		Token: token_json,
		Id:    id,
	}
	resp, err := s.Spotify_Service.GetTrackAudioAnalysis(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp.AudioAnalysis)
	fmt.Println("successfully finished writing audio analysis")
}

func (s *SpotifyApi) GetTrackAudioFeatures(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	token, err := GetToken(s, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token_json, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req := spotifyv1.GetTrackAudioFeaturesRequest{
		Token: token_json,
		Id:    id,
	}
	resp, err := s.Spotify_Service.GetTrackAudioFeatures(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp.AudioFeatures)
	fmt.Println("successfully finished writing audio features")

}
