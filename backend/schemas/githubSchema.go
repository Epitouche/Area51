package schemas

import "time"

type GithubAction string

const (
	GithubPullRequest GithubAction = "pull_request"
)

type GithubReaction string

const (
	GithubReactionListComments GithubReaction = "list_comments"
)

type GitHubResponseToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type GithubUserInfo struct {
	Login     string `json:"login"`
	Id        uint64 `json:"id"         gorm:"primaryKey"`
	AvatarUrl string `json:"avatar_url"`
	Type      string `json:"type"`
	HtmlUrl   string `json:"html_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type GithubListCommentsResponse struct {
	Body           string `json:"body"`
	PullRequestUrl string `json:"pull_request_url"`
	//! Needs more fields in the future
}
type GithubPullRequestOptions struct {
	Repo      string    `json:"-"`
	Owner     string    `json:"-"`
	CheckedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}
