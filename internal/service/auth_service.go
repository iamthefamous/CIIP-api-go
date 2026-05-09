package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type userRepository interface {
	GetProfileRoleByID(userID string) (string, error)
}

type AuthService struct {
	repo            userRepository
	supabaseURL     string
	supabaseAnonKey string
	appJWTSecret    string
	httpClient      *http.Client
}

type supabaseLoginResponse struct {
	AccessToken string `json:"access_token"`
	User        struct {
		ID string `json:"id"`
	} `json:"user"`
}

func NewAuthService(repo userRepository, supabaseURL, supabaseAnonKey, appJWTSecret string) *AuthService {
	return &AuthService{
		repo:            repo,
		supabaseURL:     strings.TrimRight(supabaseURL, "/"),
		supabaseAnonKey: supabaseAnonKey,
		appJWTSecret:    appJWTSecret,
		httpClient:      http.DefaultClient,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	if s.supabaseURL == "" || s.supabaseAnonKey == "" || s.appJWTSecret == "" {
		return "", errors.New("supabase auth config missing")
	}

	body, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, s.supabaseURL+"/auth/v1/token?grant_type=password", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("apikey", s.supabaseAnonKey)
	req.Header.Set("Authorization", "Bearer "+s.supabaseAnonKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid credentials")
	}

	var loginResp supabaseLoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return "", err
	}

	if loginResp.AccessToken == "" || loginResp.User.ID == "" {
		return "", errors.New("invalid auth response")
	}

	role, err := s.repo.GetProfileRoleByID(loginResp.User.ID)
	if err != nil || role != "admin" {
		return "", errors.New("invalid credentials")
	}

	appToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": loginResp.User.ID,
		"role":    "admin",
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	})

	signedToken, err := appToken.SignedString([]byte(s.appJWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
