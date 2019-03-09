package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	GithubPAT string
	Repositories []Repository
}

type Repository struct {
	Type string
	Organization string
	Name string
	FullName string
}


func Get() (config Config) {
	pat := os.Getenv("GITHUB_PAT")
	if pat == "" {
		panic(errors.New("No Github PAT passed."))
	}
	config.GithubPAT = pat

	i := 1
	for i != -1 {
		var envText strings.Builder
		envText.WriteString("GIT_REPO")
		envText.WriteString(strconv.Itoa(i))
		repo := os.Getenv(envText.String())

		if repo == "" {
			i = -1
			break
		}

		splitedRepo := strings.Split(repo, "/")

		repository := Repository{
			FullName: repo,
			Name: splitedRepo[1],
			Organization: splitedRepo[0],
			Type: "github",
		}

		config.Repositories = append(config.Repositories, repository)
		i++
	}
	return
}