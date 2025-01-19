package schemas

import "errors"

type ServicesUserInfos struct {
	GithubUserInfos    *GithubUserInfo    `json:"github_user_infos"`
	SpotifyUserInfos   *SpotifyUserInfo   `json:"spotify_user_infos"`
	GoogleUserInfos    *GoogleUserInfo    `json:"google_user_infos"`
	MicrosoftUserInfos *MicrosoftUserInfo `json:"microsoft_user_infos"`
}

type MobileUsefulInfos struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

var (
	ErrorNoServiceFound = errors.New("no service found")
	ErrorInvalidToken   = errors.New("invalid token")
)
