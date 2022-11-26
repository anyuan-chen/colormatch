package api_auth

import (
	"github.com/anyuan-chen/colormatch/pkg/session_management_service"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type AuthAPI struct {
	SpotifyAuthenticator spotifyauth.Authenticator
	session_manager      session_management_service.Session_Management_Service
}
