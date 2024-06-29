// This file is to create a callback API for Google sign-in
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/doptime/doptime/api"
)

type ReqSignInGoogle struct {
	ClientID     string
	ClientSecret string
	Code         string
	RedirectURI  string
	State        string
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type GoogleAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
	Error        string `json:"error"`
	ErrorDesc    string `json:"error_description"`
	GoogleUser   *GoogleUser
}

var ApiSignInGoogle = api.Api(func(req *ReqSignInGoogle) (tokenResponse *GoogleAccessTokenResponse, err error) {
	if req.Code == "" {
		return nil, errors.New("authorization code cannot be empty")
	}
	// Construct the request to exchange the code for an access token
	data := map[string]string{
		"client_id":     req.ClientID,
		"client_secret": req.ClientSecret,
		"code":          req.Code,
		"redirect_uri":  req.RedirectURI,
		"grant_type":    "authorization_code",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %v", err)
	}

	resp, err := http.Post("https://oauth2.googleapis.com/token", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Google: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to exchange authorization code for access token: %v", resp.Status)
	}

	tokenResponse = &GoogleAccessTokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	if tokenResponse.Error != "" {
		log.Printf("Google OAuth error: %v - %v", tokenResponse.Error, tokenResponse.ErrorDesc)
		return nil, fmt.Errorf("google OAuth error: %v", tokenResponse.ErrorDesc)
	}
	tokenResponse.GoogleUser, err = getGoogleUserInfo(tokenResponse.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get Google user info: %v", err)
	}

	return tokenResponse, nil
}).Func

func getGoogleUserInfo(accessToken string) (*GoogleUser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo?alt=json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: %v", resp.Status)
	}

	var user GoogleUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	return &user, nil
}
