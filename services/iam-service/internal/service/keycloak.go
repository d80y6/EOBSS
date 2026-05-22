package service

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
)

type IdentityService interface {
	Login(ctx context.Context, username, password string) (*gocloak.JWT, error)
	ValidateToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error)
	CreateUser(ctx context.Context, user gocloak.User) (string, error)
}

type keycloakService struct {
	client       *gocloak.GoCloak
	clientId     string
	clientSecret string
	realm        string
}

func NewKeycloakService(client *gocloak.GoCloak, clientId, clientSecret, realm string) IdentityService {
	return &keycloakService{
		client:       client,
		clientId:     clientId,
		clientSecret: clientSecret,
		realm:        realm,
	}
}

func (s *keycloakService) Login(ctx context.Context, username, password string) (*gocloak.JWT, error) {
	return s.client.Login(ctx, s.clientId, s.clientSecret, s.realm, username, password)
}

func (s *keycloakService) ValidateToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error) {
	return s.client.RetrospectToken(ctx, token, s.clientId, s.clientSecret, s.realm)
}

func (s *keycloakService) CreateUser(ctx context.Context, user gocloak.User) (string, error) {
	token, err := s.client.LoginClient(ctx, s.clientId, s.clientSecret, s.realm)
	if err != nil {
		return "", err
	}
	return s.client.CreateUser(ctx, token.AccessToken, s.realm, user)
}
