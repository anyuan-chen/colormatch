package spotify_service

import (
	"context"
	"encoding/json"

	imagesv1 "github.com/anyuan-chen/colormatch/gen/proto/go/images/v1"
	sharedv1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type SpotifyRetrievalServer struct {
	Auth           spotify.Authenticator
	Images_Service imagesv1.PaletteMatchingServiceClient
	spotifyv1.UnimplementedSpotifyImageColorMatchingServiceServer
}

func (s *SpotifyRetrievalServer) GetClient(token *oauth2.Token) *spotify.Client {
	spotify_client := s.Auth.NewClient(token)
	return &spotify_client
}

func (s *SpotifyRetrievalServer) GetColorMetadataForSpotifyAsset(ctx context.Context, req *spotifyv1.GetColorMetadataForSpotifyAssetRequest) (*spotifyv1.GetColorMetadataForSpotifyAssetResponse, error) {
	var token *oauth2.Token
	json.Unmarshal(req.Token, token)
	client := s.GetClient(token)
	song_title := req.SongTitle
	background_colors := req.BackgroundColors
	highlight_colors := req.HighlightColors
	overall_palette := append(background_colors.Color, highlight_colors.Color...)
	var search_type spotify.SearchType
	req_search_type := req.SearchType
	//convert search protobuf type into ones usable by Spotify
	switch req_search_type {
	case spotifyv1.SearchType_SEARCH_TYPE_ALBUM:
		search_type = spotify.SearchTypeAlbum
	case spotifyv1.SearchType_SEARCH_TYPE_ARTIST:
		search_type = spotify.SearchTypeArtist
	}
	result, err := client.Search(song_title, search_type)
	if err != nil {
		return &spotifyv1.GetColorMetadataForSpotifyAssetResponse{}, err
	}
	image_url, err := GetImageUrl(req_search_type, *result)
	if err != nil {
		return nil, err
	}
	image, err := ImageURLToImage(image_url)
	if err != nil {
		return nil, err
	}
	width := image.Bounds().Max.X - image.Bounds().Min.X
	height := image.Bounds().Max.Y - image.Bounds().Min.Y
	request_image, err := ImageToBytes(image)
	if err != nil {
		return nil, err
	}
	background_color_request := &imagesv1.GetBackgroundColorRequest{
		Image: &sharedv1.Image{
			Width:     int32(width),
			Height:    int32(height),
			ImageData: request_image,
		},
		BackgroundColors: background_colors,
		Palette:          &sharedv1.Palette{Color: overall_palette},
	}
	background_res, err := s.Images_Service.GetBackgroundColor(context.Background(), background_color_request)
	if err != nil {
		return nil, err
	}

	highlight_color_request := &imagesv1.GetHighlightColorRequest{
		Image: &sharedv1.Image{
			Width:     int32(width),
			Height:    int32(height),
			ImageData: request_image,
		},
		HighlightColors: highlight_colors,
		Palette:         &sharedv1.Palette{Color: overall_palette},
	}
	highlight_res, err := s.Images_Service.GetHighlightColor(context.Background(), highlight_color_request)
	if err != nil {
		return nil, err
	}
	response := &spotifyv1.GetColorMetadataForSpotifyAssetResponse{
		Response: &spotifyv1.ColorMetadataResponse{
			BackgroundColor: background_res.Color,
			HighlightColor:  highlight_res.Color,
		},
	}
	return response, nil
}
