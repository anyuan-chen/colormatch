package api_auth

import (
	session_managementv1 "github.com/anyuan-chen/colormatch/gen/proto/go/session_management/v1"
)

type AuthAPI struct {
	Session_Manager      session_managementv1.SessionManagementServiceClient
}
