package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AmitKarnam/WorkoutTracker/internal/config"
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuthService interface {
	GetGoogleOAuthURL(state string) string
	HandleGoogleCallback(ctx context.Context, code string) (*models.User, error)
}

type googleAuthService struct {
	userRepo repository.UserRepository
	config   *oauth2.Config
}

func NewGoogleAuthService(userRepo repository.UserRepository) GoogleAuthService {
	googleConfig := config.GetGoogleOAuthConfig()
	return &googleAuthService{
		userRepo: userRepo,
		config: &oauth2.Config{
			ClientID:     googleConfig.ClientID,
			ClientSecret: googleConfig.ClientSecret,
			RedirectURL:  googleConfig.RedirectURL,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (s *googleAuthService) GetGoogleOAuthURL(state string) string {
	return s.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s *googleAuthService) HandleGoogleCallback(ctx context.Context, code string) (*models.User, error) {
	token, err := s.config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %v", err)
	}

	client := s.config.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response from userinfo endpoint: %s", resp.Status)
	}

	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	user := &models.User{
		Email: userInfo.Email,
		Name:  userInfo.Name,
		Role:  models.Customer,
	}
	userInstance, err := s.userRepo.FindOrCreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return userInstance, nil
}
