package github

import (
	"context"
	"github.com/google/go-github/v24/github"
	"gitops/backend/config"
	"golang.org/x/oauth2"
)

func GetBranches(owner string, repo string) (branches []*github.Branch, err error) {
	ctx := context.Background()
	client := getClient()
	branches, _, err = client.Repositories.ListBranches(ctx, owner, repo, &github.ListOptions{PerPage: 100})

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