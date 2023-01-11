package spotify_api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
)

func (s *SpotifyApi) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	artistIds := r.Form["artist_ids"]
	trackIds := r.Form["track_ids"]
	genres := r.Form["genres"]
	recommendationCount, err := strconv.Atoi(r.Form.Get("recommendation_count"))
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
	encoded_resp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(encoded_resp)
}

func (s *SpotifyApi) GetTrackAudioAnalysis(w http.ResponseWriter, r *http.Request) {
	id := r.Form.Get("id")
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
	encoded_resp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(encoded_resp)
}

func (s *SpotifyApi) GetTrackAudioFeatures(w http.ResponseWriter, r *http.Request) {
	id := r.Form.Get("id")
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
	encoded_resp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(encoded_resp)
}
