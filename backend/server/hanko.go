package server

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SessionValidator interface {
	ValidateSession(token string) (bool, error)
}

type HankoSessionValidator struct {
	apiURL string
}

type ValidationResponse struct {
	IsValid bool `json:"is_valid"`
}

func NewHankoSessionValidator(apiURL string) *HankoSessionValidator {
	return &HankoSessionValidator{apiURL: apiURL}
}

func (v *HankoSessionValidator) ValidateSession(token string) (bool, error) {
	payload := strings.NewReader(fmt.Sprintf(`{"session_token":"%s"}`, token))

	req, err := http.NewRequest(http.MethodPost, v.apiURL+"/sessions/validate", payload)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response: %w", err)
	}

	var validationRes ValidationResponse
	if err := json.Unmarshal(body, &validationRes); err != nil {
		return false, fmt.Errorf("failed to parse response: %w", err)
	}

	return validationRes.IsValid, nil
}

func AuthMiddleware(validator SessionValidator) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-API-KEY") != "" {
				//
				//HANDLE API KEYS HERE
				//
			} else {
				cookie, err := r.Cookie("hanko")
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				isValid, err := validator.ValidateSession(cookie.Value)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				if !isValid {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

			}
			next(w, r)
		})
	}
}

func GetUserFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("hanko")
	if err != nil {
		return "", errors.New("auth cookie does not exist")
	}
	jwtParts := strings.Split(cookie.Value, ".")

	if len(jwtParts) < 2 {
		fmt.Println("Invalid token")
		return "", errors.New("invalid login token")
	}

	data, err := base64.RawURLEncoding.DecodeString(jwtParts[1])
	if err != nil {
		return "", errors.New("failed to decode JWT claim")
	}

	var pretty JWTClaims
	if err := json.Unmarshal(data, &pretty); err != nil {
		return "", errors.New("failed to parse JWT claim JSON")
	}

	if pretty.Subject != "" {
		return pretty.Subject, nil
	} else {
		return "", errors.New("failed to retrieve email from JWT")
	}
}

type JWTClaims struct {
	Audience []string `json:"aud"`
	Email    struct {
		Address    string `json:"address"`
		IsPrimary  bool   `json:"is_primary"`
		IsVerified bool   `json:"is_verified"`
	} `json:"email"`
	ExpiresAt int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
	Issuer    string `json:"iss"`
	SessionID string `json:"session_id"`
	Subject   string `json:"sub"`
}
