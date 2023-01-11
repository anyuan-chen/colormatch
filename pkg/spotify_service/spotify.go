package spotify_service

import (
	imagesv1 "github.com/anyuan-chen/colormatch/gen/proto/go/images/v1"
	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	"github.com/zmb3/spotify"
)

type SpotifyRetrievalServer struct {
	Auth           spotify.Authenticator
	Images_Service imagesv1.PaletteMatchingServiceClient
	spotifyv1.UnimplementedSpotifyAPIServiceServer
}
