package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type GithubRepository interface {
	Save(githubPullRequestOptions schemas.GithubPullRequestOptionsTable)
	Update(githubPullRequestOptions schemas.GithubPullRequestOptionsTable)
	Delete(githubPullRequestOptions schemas.GithubPullRequestOptionsTable)
	UpdateNumPRs(githubPullRequestOptions schemas.GithubPullRequestOptionsTable)

	FindAll() []schemas.GithubPullRequestOptionsTable
	FindByOwnerAndRepo(owner string, repository string) schemas.GithubPullRequestOptionsTable
	// FindByWorkflowId(workflowId uint64)
}

type githubRepository struct {
	db *schemas.Database
}

func NewGithubRepository(db *gorm.DB) GithubRepository {
	err := db.AutoMigrate(&schemas.GithubPullRequestOptionsTable{})

	if err != nil {
		panic("failed to migrate database")
	}
	return &githubRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}

func (repo *githubRepository) Save(githubPullRequestOptions schemas.GithubPullRequestOptionsTable) {
	err := repo.db.Connection.Create(&githubPullRequestOptions)

	if err.Error != nil {
		return
	}
}

func (repo *githubRepository) Update(githubPullRequestOptions schemas.GithubPullRequestOptionsTable) {
	err := repo.db.Connection.Where(&schemas.GithubPullRequestOptionsTable{
		Id: githubPullRequestOptions.Id,
	}).Updates(&githubPullRequestOptions)

	if err.Error != nil {
		return
	}
}

func (repo *githubRepository) Delete(githubPullRequestOptions schemas.GithubPullRequestOptionsTable) {
	err := repo.db.Connection.Delete(&githubPullRequestOptions)

	if err.Error != nil {
		return
	}
}

func (repo *githubRepository) FindAll() (githubPullRequestOptions []schemas.GithubPullRequestOptionsTable) {
	err := repo.db.Connection.Find(&githubPullRequestOptions)

	if err.Error != nil {
		return []schemas.GithubPullRequestOptionsTable{}
	}
	return githubPullRequestOptions
}

func (repo *githubRepository) FindByOwnerAndRepo(owner string, repository string) (githubPullRequestOptions schemas.GithubPullRequestOptionsTable) {
	err := repo.db.Connection.Where(&schemas.GithubPullRequestOptionsTable{
		Owner: owner,
		Repo:  repository,
	}).First(&githubPullRequestOptions)

	if err.Error != nil {
		return schemas.GithubPullRequestOptionsTable{}
	}
	return githubPullRequestOptions
}

func (repo *githubRepository) UpdateNumPRs(githubPullRequestOptions schemas.GithubPullRequestOptionsTable) {
	err := repo.db.Connection.Model(&schemas.GithubPullRequestOptionsTable{}).Where(&schemas.GithubPullRequestOptionsTable{Id: githubPullRequestOptions.Id}).Updates(map[string]interface{}{
		"num_pr": githubPullRequestOptions.NumPR,
	})
	if err.Error != nil {
		return
	}
}
