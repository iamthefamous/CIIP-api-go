package service

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockUserRepository struct {
	getProfileRoleByIDFn func(userID string) (string, error)
}

func (m *mockUserRepository) GetProfileRoleByID(userID string) (string, error) {
	if m.getProfileRoleByIDFn == nil {
		return "", nil
	}
	return m.getProfileRoleByIDFn(userID)
}

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newMockHTTPClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestAuthServiceLoginSuccess(t *testing.T) {
	svc := NewAuthService(&mockUserRepository{
		getProfileRoleByIDFn: func(userID string) (string, error) {
			if userID != "u1" {
				t.Fatalf("expected user id u1, got %s", userID)
			}
			return "admin", nil
		},
	}, "https://example.supabase.co", "anon-key")
	svc.httpClient = newMockHTTPClient(func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/auth/v1/token" || req.URL.RawQuery != "grant_type=password" {
			t.Fatalf("unexpected auth path: %s?%s", req.URL.Path, req.URL.RawQuery)
		}
		if got := req.Header.Get("apikey"); got != "anon-key" {
			t.Fatalf("expected apikey header anon-key, got %s", got)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"access_token":"supabase-token","user":{"id":"u1"}}`)),
			Header:     make(http.Header),
		}, nil
	})

	token, err := svc.Login("admin@example.com", "secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "supabase-token" {
		t.Fatalf("expected supabase-token, got %s", token)
	}
}

func TestAuthServiceLoginInvalidCredentialsFromSupabase(t *testing.T) {
	svc := NewAuthService(&mockUserRepository{}, "https://example.supabase.co", "anon-key")
	svc.httpClient = newMockHTTPClient(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       io.NopCloser(strings.NewReader(`{"error":"invalid_grant"}`)),
			Header:     make(http.Header),
		}, nil
	})

	token, err := svc.Login("admin@example.com", "bad-password")
	if err == nil {
		t.Fatal("expected error")
	}
	if token != "" {
		t.Fatalf("expected empty token, got %s", token)
	}
}

func TestAuthServiceLoginForbiddenForNonAdminProfile(t *testing.T) {
	svc := NewAuthService(&mockUserRepository{
		getProfileRoleByIDFn: func(userID string) (string, error) {
			return "user", nil
		},
	}, "https://example.supabase.co", "anon-key")
	svc.httpClient = newMockHTTPClient(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"access_token":"supabase-token","user":{"id":"u1"}}`)),
			Header:     make(http.Header),
		}, nil
	})

	_, err := svc.Login("admin@example.com", "secret")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestAuthServiceLoginProfileLookupError(t *testing.T) {
	repoErr := errors.New("query failed")
	svc := NewAuthService(&mockUserRepository{
		getProfileRoleByIDFn: func(userID string) (string, error) {
			return "", repoErr
		},
	}, "https://example.supabase.co", "anon-key")
	svc.httpClient = newMockHTTPClient(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"access_token":"supabase-token","user":{"id":"u1"}}`)),
			Header:     make(http.Header),
		}, nil
	})

	_, err := svc.Login("admin@example.com", "secret")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestAuthServiceLoginMissingConfig(t *testing.T) {
	svc := NewAuthService(&mockUserRepository{}, "", "")

	_, err := svc.Login("admin@example.com", "secret")
	if err == nil {
		t.Fatal("expected error")
	}
}
