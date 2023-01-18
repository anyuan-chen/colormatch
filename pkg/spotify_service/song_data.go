package spotify_service

import (
	"context"
	"encoding/json"
	"fmt"

	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func (s *SpotifyRetrievalServer) GetTrackAudioFeatures(ctx context.Context, req *spotifyv1.GetTrackAudioFeaturesRequest) (*spotifyv1.GetTrackAudioFeaturesResponse, error) {
	var token oauth2.Token
	json.Unmarshal(req.Token, &token)
	client := s.GetClient(&token)
	id := spotify.ID(req.Id)
	fmt.Println("cleared features checks")
	features, err := client.GetAudioFeatures(id)
	fmt.Println("done features checks")

	if err != nil {
		return &spotifyv1.GetTrackAudioFeaturesResponse{}, err
	}
	encoded_features, err := json.Marshal(features)
	if err != nil {
		return nil, err
	}
	fmt.Println("sending features ")
	return &spotifyv1.GetTrackAudioFeaturesResponse{
		AudioFeatures: encoded_features,
	}, nil
}
func (s *SpotifyRetrievalServer) GetTrackAudioAnalysis(ctx context.Context, req *spotifyv1.GetTrackAudioAnalysisRequest) (*spotifyv1.GetTrackAudioAnalysisResponse, error) {
	var token oauth2.Token
	json.Unmarshal(req.Token, &token)
	client := s.GetClient(&token)
	id := spotify.ID(req.Id)
	fmt.Println("cleared analysis checks")
	analysis, err := client.GetAudioAnalysis(id)
	fmt.Println("done analysis retrieval")
	if err != nil {
		return &spotifyv1.GetTrackAudioAnalysisResponse{}, err
	}
	encoded_analysis, err := json.Marshal(analysis)
	if err != nil {
		return nil, err
	}
	fmt.Println("sending analysis")
	return &spotifyv1.GetTrackAudioAnalysisResponse{
		AudioAnalysis: encoded_analysis,
	}, nil
}
func (s *SpotifyRetrievalServer) GetRecommendations(ctx context.Context, req *spotifyv1.GetRecommendationsRequest) (*spotifyv1.GetRecommendationsResponse, error) {
	var token oauth2.Token
	json.Unmarshal(req.Token, &token)
	client := s.GetClient(&token)
	recommendation_count := int(req.RecommendationCount)
	artist_ids := make([]spotify.ID, 0, len(req.ArtistIds))
	for i := 0; i < len(req.ArtistIds); i++ {
		artist_ids = append(artist_ids, spotify.ID(req.ArtistIds[i]))
	}
	track_ids := make([]spotify.ID, 0, len(req.TrackIds))
	for i := 0; i < len(req.TrackIds); i++ {
		track_ids = append(track_ids, spotify.ID(req.TrackIds[i]))
	}
	seeds := spotify.Seeds{
		Artists: artist_ids,
		Tracks:  track_ids,
		Genres:  req.Genres,
	}
	track_attributes := spotify.TrackAttributes{}
	recommendations, err := client.GetRecommendations(seeds, &track_attributes, &spotify.Options{
		Limit: &recommendation_count,
	})
	if err != nil {
		return nil, err
	}
	encoded_recommendations, err := json.Marshal(recommendations)
	if err != nil {
		return nil, err
	}
	return &spotifyv1.GetRecommendationsResponse{
		Recommendations: encoded_recommendations,
	}, nil
}
