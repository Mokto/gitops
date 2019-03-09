package github

import (
	"errors"
	"fmt"
	"gitops/backend/config"
	"os"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

// CloneRepo clones a repo and reset it
func CloneRepo() (repo *git.Repository, err error) {
	path := "/tmp/cloned-repo"

	appConfig := config.Get()
	if appConfig .GithubPAT == "" {
		return nil, errors.New("No Github PAT passed")
	}

	var URL strings.Builder
	URL.WriteString("https://")
	URL.WriteString(appConfig .GithubPAT)
	URL.WriteString(":x-oauth-basic@github.com/")
	URL.WriteString(appConfig .Repositories[0].FullName)
	URL.WriteString(".git")

	repo, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:      URL.String(),
		Progress: os.Stdout,
	})
	fmt.Println("Plain clone repo...")
	if err != nil && err.Error() != "repository already exists" {
		return nil, err
	}
	if err != nil {
		repo, err = git.PlainOpen(path)
		if err != nil {
			return nil, err
		}
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}
	worktree.Pull(&git.PullOptions{})
	worktree.Reset(&git.ResetOptions{Mode: git.HardReset})
	return repo, nil
}
