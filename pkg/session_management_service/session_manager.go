package session_management_service

import (
	"encoding/json"

	session_managementv1 "github.com/anyuan-chen/colormatch/gen/proto/go/session_management/v1"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type Session_Management_Service struct {
	Sessions map[string]*oauth2.Token
	session_managementv1.UnimplementedSessionManagementServiceServer
}

func (s *Session_Management_Service) Retrieve(ctx context.Context, r *session_managementv1.RetrieveRequest) (*session_managementv1.RetrieveResponse, error) {
	ciphertext := r.Ciphertext
	json_token, err := json.Marshal(s.Sessions[ciphertext])
	if err != nil {
		return nil, err
	}
	return &session_managementv1.RetrieveResponse{Token: json_token}, nil
}

func (s *Session_Management_Service) SetToken(ctx context.Context, r *session_managementv1.SetTokenRequest) (*session_managementv1.SetTokenResponse, error) {
	var token *oauth2.Token
	json.Unmarshal(r.Token, token)
	id := uuid.New()
	for s.Sessions[id.String()] != nil {
		id = uuid.New()
	}
	return &session_managementv1.SetTokenResponse{Ciphertext: id.String()}, nil
}
