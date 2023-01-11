package spotify_service

import (
	"context"
	"encoding/json"
	"fmt"

	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func (s *SpotifyRetrievalServer) GetClient(token *oauth2.Token) *spotify.Client {
	spotify_client := spotify.Authenticator.NewClient(s.Auth, token)
	return &spotify_client
}
func (s *SpotifyRetrievalServer) PingTokenValidity(ctx context.Context, req *spotifyv1.PingTokenValidityRequest) (*spotifyv1.PingTokenValidityResponse, error) {
	var token oauth2.Token
	json.Unmarshal(req.Token, &token)
	fmt.Println("token recieved by spotify: ", token)
	client := s.GetClient(&token)
	fmt.Println("client retrieved successfully for ping token validity")
	_, err := client.Search("jay chou", spotify.SearchTypeArtist)
	fmt.Println("ping spotify search complete")
	if err != nil {
		fmt.Println("problem with the request" + err.Error())
		return &spotifyv1.PingTokenValidityResponse{Valid: false}, nil
	}
	return &spotifyv1.PingTokenValidityResponse{Valid: true}, nil
}
