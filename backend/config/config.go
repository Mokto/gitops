package config

import (
	"errors"
	"gitops/backend/models"
	"gitops/backend/utils"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	GithubPAT string
}

func Get() (config Config) {
	pat := os.Getenv("GITHUB_PAT")
	if pat == "" {
		panic(errors.New("No Github PAT passed."))
	}
	config.GithubPAT = pat
	return
}

func GetRepositories() (repositories []models.Repository) {
	i := 1
	for i != -1 {
		envText := utils.ComposeStrings("GIT_REPO", strconv.Itoa(i))
		repo := os.Getenv(envText)

		if repo == "" {
			i = -1
			break
		}

		splitedRepo := strings.Split(repo, "/")

		repository := models.Repository{
			FullName: repo,
			Name: splitedRepo[1],
			Organization: splitedRepo[0],
			Type: "github",
		}

		repositories = append(repositories, repository)
		i++
	}
	return repositories
}