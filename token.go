package peg_go_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GetAuthTokenRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type GetAuthTokenErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
	} `json:"error"`
}

type GetAuthTokenSuccessResponse struct {
	Kind         string `json:"kind"`
	LocalId      string `json:"localId"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	IdToken      string `json:"idToken"`
	Registered   bool   `json:"registered"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func GetAuthToken(email string, password string) (string, string, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", os.Getenv("FB_API_KEY"))

	req := GetAuthTokenRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return "", "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var r GetAuthTokenErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			return "", "", err
		}
		return "", "", fmt.Errorf("failed to get auth token(code=%v): %v", r.Error.Code, r.Error.Message)
	} else {
		var r GetAuthTokenSuccessResponse
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			return "", "", err
		}

		return r.IdToken, r.LocalId, nil
	}
}
