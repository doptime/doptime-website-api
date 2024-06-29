// this file is to create a callback api for github sign in
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/doptime/doptime/api"
	"github.com/go-resty/resty/v2"
)

type ReqSignInGithub struct {
	ClientID     string
	ClientSecret string
	Code         string
	RedirectURI  string
	State        string
}
type GitHubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Error       string `json:"error"`
	ErrorDesc   string `json:"error_description"`
}

var ApiSignInGithub = api.Api(func(req *ReqSignInGithub) (githubUser *GitHubUser, err error) {
	if req.Code == "" {
		return nil, errors.New("authorization code cannot be empty")
	}
	// Construct the request to exchange the code for an access token
	data := map[string]string{
		"client_id":     req.ClientID,
		"client_secret": req.ClientSecret,
		"code":          req.Code,
		"redirect_uri":  req.RedirectURI,
		"state":         req.State,
	}
	var tokenResponse GitHubAccessTokenResponse
	resp, err := resty.New().R().SetHeader("Content-Type", "application/json").SetBody(data).SetResult(&tokenResponse).Post("https://github.com/login/oauth/access_token")
	if err != nil {
		return nil, err
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to exchange authorization code for access token")
	}
	if tokenResponse.Error != "" {
		return nil, errors.New("GitHub OAuth error: " + tokenResponse.ErrorDesc)
	} else if tokenResponse.AccessToken == "" {
		return nil, errors.New("GitHub OAuth error: access_token not found")
	}
	return getUserInfo(tokenResponse.AccessToken)
}).Func

type GitHubUser struct {
	Login     string `json:"login"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func getUserInfo(accessToken string) (*GitHubUser, error) {

	var user GitHubUser
	client := resty.New()
	// 获取用户基本信息
	resp, err := client.R().SetHeader("Authorization", "Bearer "+accessToken).SetResult(&user).Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	} else if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch user info: %s", resp.Status())
	}
	if user.Email == "" {
		var emails []struct {
			Email      string `json:"email"`
			Primary    bool   `json:"primary"`
			Verified   bool   `json:"verified"`
			Visibility string `json:"visibility"`
		}
		resp, err = client.R().SetHeader("Authorization", "Bearer "+accessToken).SetResult(&emails).Get("https://api.github.com/user/emails")
		if err != nil {
			return nil, err
		} else if resp.StatusCode() != 200 {
			return nil, fmt.Errorf("failed to fetch user emails: %s", resp.Status())
		}
		for _, email := range emails {
			if email.Primary && email.Verified {
				user.Email = email.Email
				break
			}
		}
	}

	return &user, nil
}
