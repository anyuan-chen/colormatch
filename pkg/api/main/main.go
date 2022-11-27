package main

import (
	"log"
	"net/http"
	"os"

	session_managementv1 "github.com/anyuan-chen/colormatch/gen/proto/go/session_management/v1"
	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	api_auth "github.com/anyuan-chen/colormatch/pkg/api/auth_api"
	spotify_api "github.com/anyuan-chen/colormatch/pkg/api/spotify_api"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type Server struct {
	r *mux.Router
}

func main() {
	r := mux.NewRouter()
	//session manager
	session_management_connection, err := grpc.Dial(os.Getenv("SESSION_MANAGEMENT_PORT"))
	if err != nil {
		log.Fatal("failed to connect to session_management_service")
	}
	defer session_management_connection.Close()
	session_management_service_client := session_managementv1.NewSessionManagementServiceClient(session_management_connection)
	//spotify
	spotify_service_connection, err := grpc.Dial(os.Getenv("SPOTIFY_SERVICE_PORT"))
	if err != nil {
		log.Fatal("failed to connect to spotify_service")
	}
	defer spotify_service_connection.Close()
	spotify_service_client := spotifyv1.NewSpotifyImageColorMatchingServiceClient(spotify_service_connection)
	//auth
	auth := r.PathPrefix("/auth").Subrouter()
	authAPI := &api_auth.AuthAPI{}
	auth.HandleFunc("/spotify/login", authAPI.Login)
	auth.HandleFunc("/spotify/callback", authAPI.Callback)

	spotify := r.PathPrefix("/spotify").Subrouter()
	spotifyAPI := &spotify_api.SpotifyApi{
		Session_Manager: session_management_service_client,
		Spotify_Service: spotify_service_client,
	}
	spotify.HandleFunc("/colors", spotifyAPI.GetColorSummary)
	spotify.HandleFunc("/ping", spotifyAPI.PingTokenValidity)
	http.Handle("/", &Server{r: r})
	http.ListenAndServe(":8080", nil)
}
func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if req.Method == "OPTIONS" {
		return
	}
	s.r.ServeHTTP(rw, req)
}
