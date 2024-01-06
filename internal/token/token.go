package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Name  string   `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

func ValidateTokenWithIdp(token string, identity_provider_endpoint string) (*JwtCustomClaims, error) {
	h := http.Client{}
	req, err := http.NewRequest("GET", identity_provider_endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	r, err := h.Do(req)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, errors.New("invalid token")
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	claims := JwtCustomClaims{}
	if err := json.Unmarshal(body, &claims); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return &claims, nil
}
