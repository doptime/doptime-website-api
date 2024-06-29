// https: //api.doptime.com/API-!SignInGithubToDoptime
// this file is to create a callback api for github sign in
package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/doptime/doptime/api"
	"github.com/doptime/doptime/config"
	"github.com/doptime/doptime/libapi"
	"github.com/doptime/doptime/rdsdb"
)

type ReqSignInGithubToDoptime struct {
	Code        string
	RedirectURI string
	State       string
}
type DoptimeUser struct {
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatar_url"`
	GithubID    int64  `json:"-"`
	ID          string `json:"id"`
	AccessToken string `json:"jwt" msgpack:"-"`
}

var keyDoptimeUser = rdsdb.HashKey[string, *DoptimeUser]()

var APISignInGithubToDoptime = api.Api(func(req *ReqSignInGithubToDoptime) (user *DoptimeUser, err error) {
	var AccessToken string
	ClientSecret, ClientID := os.Getenv("GITHUB_CLIENT_SECRET"), os.Getenv("GITHUB_CLIENT_ID")
	if ClientSecret == "" {
		return nil, errors.New("GITHUB_CLIENT_SECRET is not set")
	} else if config.Cfg.Http.JwtSecret == "" {
		return nil, errors.New("JWT_SECRET is not set")
	}
	GitHubUser, err := ApiSignInGithub(&ReqSignInGithub{ClientID: ClientID, ClientSecret: ClientSecret, Code: req.Code, RedirectURI: req.RedirectURI, State: req.State})
	if err != nil {
		return nil, err
	}
	user = &DoptimeUser{
		UserName:  GitHubUser.Name,
		Email:     GitHubUser.Email,
		AvatarURL: GitHubUser.AvatarURL,
		GithubID:  GitHubUser.ID,
		ID:        "github_" + strconv.Itoa(int(GitHubUser.ID)),
	}
	signData := map[string]interface{}{"UserName": user.UserName, "Email": user.Email, "AvatarUrl": user.AvatarURL, "ID": user.ID}

	user.AccessToken, err = libapi.ApiJwtSign(&libapi.JwtEncodingIn{Params: signData, JwtSecret: config.Cfg.Http.JwtSecret, SignMethod: "HS256", Duration: 3600 * 24 * 30 * 12})
	keyDoptimeUser.HSet(user.Email, user)
	signData["jwt"] = AccessToken
	return user, err
}).Func
