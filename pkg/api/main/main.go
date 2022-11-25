package main

import (
	"net/http"

	api_auth "github.com/anyuan-chen/colormatch/pkg/api/auth_api"
	spotify_api "github.com/anyuan-chen/colormatch/pkg/api/spotify_api"
	"github.com/gorilla/mux"
)

type Server struct {
	r *mux.Router
}

func main() {
	r := mux.NewRouter()

	//auth
	auth := r.PathPrefix("/auth").Subrouter()
	authAPI := &api_auth.AuthAPI{}
	auth.HandleFunc("/spotify/login", authAPI.Login)
	auth.HandleFunc("/spotify/callback", authAPI.Callback)

	spotify := r.PathPrefix("/spotify").Subrouter()
	spotifyAPI := &spotify_api.SpotifyApi{}
	spotify.HandleFunc("/artist", spotifyAPI.GetColorSummaryForArtist)
	spotify.HandleFunc("/album", spotifyAPI.GetColorSummaryForAlbum)

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
