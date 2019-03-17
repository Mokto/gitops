package github

import (
	"context"
	"github.com/google/go-github/v24/github"
	"gitops/backend/config"
	"gitops/backend/models"
	"golang.org/x/oauth2"
)

func GetBranches(owner string, repo string) (branches []models.Branch, err error) {
	ctx := context.Background()
	client := getClient()
	githubBranches, _, err := client.Repositories.ListBranches(ctx, owner, repo, &github.ListOptions{PerPage: 200})
	if err != nil {
		return nil, err
	}
	for _, githubBranch := range githubBranches {
		branches = append(branches, models.Branch{
			Name: githubBranch.GetName(),
		})
	}

	return
}

func getClient() (client *github.Client) {
	ctx := context.Background()
	appConfig := config.Get()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: appConfig.GithubPAT},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)

	return
}
