package api_auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	session_managementv1 "github.com/anyuan-chen/colormatch/gen/proto/go/session_management/v1"
	"github.com/zmb3/spotify"
)

var Authenticator = spotify.NewAuthenticator(os.Getenv("SPOTIFY_REDIRECT_URI"), spotify.ScopeUserReadPrivate)

func (a *AuthAPI) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirect uri:", os.Getenv("SPOTIFY_REDIRECT_URI"))
	fmt.Println("login endpoint hit")
	expiration := time.Now().Add(20 * time.Minute)
	state := make(map[string]interface{})
	randomizedState := make([]byte, 16)
	rand.Read(randomizedState)
	state["random"] = randomizedState
	json, err := json.Marshal(state)
	if err != nil {
		http.Error(w, "A JSON Encoding error has been encountered.", http.StatusInternalServerError)
		return
	}
	encoded_json := base64.URLEncoding.EncodeToString(json)
	fmt.Println("login: json state created " + encoded_json)
	url := Authenticator.AuthURL(encoded_json)
	fmt.Println("login: authurl created " + url)
	cookie := http.Cookie{Name: "oauthstate", Value: encoded_json, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func (a *AuthAPI) Callback(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("oauthstate")
	fmt.Println("cookie state: " + state.Value)
	fmt.Println("form value: " + r.FormValue("state"))
	if err != nil || r.FormValue("state") != state.Value {
		http.Error(w, "Bad OAuth State", http.StatusInternalServerError)
		return
	}
	token, err := Authenticator.Token(r.FormValue("state"), r)
	fmt.Println("Authenticator Token error:", err)
	if err != nil {
		http.Error(w, "Failed to Retrieve Token "+err.Error(), http.StatusInternalServerError)
		return
	}
	token_json, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Error encoding JSON token", http.StatusInternalServerError)
		return
	}
	fmt.Println("Token", token)
	set_token_request := session_managementv1.SetTokenRequest{
		Token: token_json,
	}
	set_token_response, err := a.Session_Manager.SetToken(context.Background(), &set_token_request)
	fmt.Println("Final Code: ", set_token_response.Ciphertext)
	if err != nil {
		http.Error(w, "Problem with the session management service: "+err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{Name: "spotify_session_id", Path: "/", Value: set_token_response.Ciphertext, Secure: true, Expires: time.Now().Add(time.Hour * 24 * 7)}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, os.Getenv("FRONTEND_SPOTIFY_CALLBACK_URL"), http.StatusPermanentRedirect)
}
