package api_auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	session_managementv1 "github.com/anyuan-chen/colormatch/gen/proto/go/session_management/v1"
)

func (a *AuthAPI) Login(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(20 * time.Minute)
	state := make(map[string]interface{})
	randomizedState := make([]byte, 16)
	rand.Read(randomizedState)
	state["random"] = randomizedState
	json, err := json.Marshal(state)
	if err != nil {
		http.Error(w, "A JSON Encoding error has been encountered.", http.StatusInternalServerError)
	}
	encoded_json := base64.URLEncoding.EncodeToString(json)
	url := a.SpotifyAuthenticator.AuthURL(encoded_json)
	cookie := http.Cookie{Name: "oauthstate", Value: encoded_json, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func (a *AuthAPI) Callback(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("oauthstate")
	if err != nil || r.FormValue("state") != state.Value {
		http.Error(w, "Bad OAuth State", http.StatusInternalServerError)
	}
	token, err := a.SpotifyAuthenticator.Token(r.Context(), state.Value, r)
	if err != nil {
		http.Error(w, "Failed to Retrieve Token", http.StatusInternalServerError)
	}
	token_json, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Error encoding JSON token", http.StatusInternalServerError)
	}
	set_token_request := &session_managementv1.SetTokenRequest{
		Token: token_json,
	}
	set_token_response, err := a.session_manager.SetToken(context.Background(), set_token_request)
	if err != nil {
		http.Error(w, "Problem with the session management service: "+err.Error(), http.StatusInternalServerError)
	}
	cookie := http.Cookie{Name: os.Getenv("spotify_session_id"), Value: set_token_response.Ciphertext, SameSite: http.SameSiteNoneMode, Secure: true}
	http.SetCookie(w, &cookie)
	//todo later
	http.Redirect(w, r, os.Getenv(""), http.StatusPermanentRedirect)
}
