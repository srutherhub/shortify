package server

import (
	"encoding/json"
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
		})
	}
}
