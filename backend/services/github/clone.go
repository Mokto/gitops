package github

import (
	"errors"
	"fmt"
	"gitops/backend/config"
	"gitops/backend/models"
	"gitops/backend/utils"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

// CloneRepo clones a repo and reset it
func CloneRepo(repository models.Repository, branch models.Branch) (repo *git.Repository, path string, err error) {
	path = utils.ComposeStrings("/tmp/gitops/", repository.FullName, "/", branch.Name)

	appConfig := config.Get()
	if appConfig.GithubPAT == "" {
		return nil, "", errors.New("No Github PAT passed.")
	}

	url := utils.ComposeStrings("https://", appConfig.GithubPAT, ":x-oauth-basic@github.com/", repository.FullName, ".git")

	repo, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:           url,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(branch.Name),
		SingleBranch:  true,
		Depth:         1,
	})
	fmt.Println("Plain clone repo...")
	if err != nil && err.Error() != "repository already exists" {
		return nil, "", err
	}
	if err != nil {
		repo, err = git.PlainOpen(path)
		if err != nil {
			return nil, "", err
		}
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, "", err
	}
	worktree.Pull(&git.PullOptions{})
	worktree.Reset(&git.ResetOptions{Mode: git.HardReset})
	return repo, path, nil
}
