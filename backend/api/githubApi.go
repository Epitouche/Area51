package api

import "area51/controllers"


type GithubApi struct {
	controller controllers.GithubController
}

func NewGithubApi(controller controllers.GithubController) *GithubApi {
	return &GithubApi{
		controller: controller,
	}
}
