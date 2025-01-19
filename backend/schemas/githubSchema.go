package schemas

type GithubAction string

const (
	GithubPullRequest GithubAction = "pull_request"
	GithubPushOnRepo  GithubAction = "push_on_repo"
)

type GithubReaction string

const (
	GithubReactionListComments GithubReaction = "list_review_comments"
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
	Repo  string `json:"repo"`
	Owner string `json:"owner"`
}

type GithubListAllReviewCommentsOptions struct {
	Repo  string `json:"repo"`
	Owner string `json:"owner"`
}

type GithubPushOnRepoOptions struct {
	Repo   string `json:"repo"`
	Owner  string `json:"owner"`
	Branch string `json:"branch"`
}

// type GithubPushOnRepoOptionsTable struct {
// 	Id             uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
// 	User           User      `json:"user,omitempty" gorm:"foreignkey:UserId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
// 	UserId         uint64    `json:"-"`
// 	Workflow       Workflow  `json:"workflow,omitempty" gorm:"foreignkey:WorkflowId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
// 	WorkflowId     uint64    `json:"-"`
// 	LastCommitDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_commit_date"`
// }
