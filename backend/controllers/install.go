package controllers

import (
	"fmt"
	"gitops/backend/dao"
	"gitops/backend/models"
	"gitops/backend/services/github"
	"gitops/backend/services/helm"
	"gitops/backend/services/templates"
	"gitops/backend/utils"
)

func InstallAllRepositories() (err error) {
	repositories, err := dao.FindAllRepositories()
	if err != nil {
		return err
	}

	for _, repo := range repositories {
		err = InstallRepository(repo)
		if err != nil {
			return err
		}
	}
	return
}

func InstallRepository(repository models.Repository) (err error) {
	branches, err := dao.FindAllBranchesByRepository(repository.ID)
	if err != nil {
		return err
	}

	for _, branch := range branches {
		err = InstallRepositoryBranch(repository, branch)
		if err != nil {
			return err
		}
	}

	return
}

func InstallRepositoryBranch(repository models.Repository, branch models.Branch) (err error) {
	_, path, err := github.CloneRepo(repository, branch)
	if err != nil {
		fmt.Println(err)
		return err
	}

	secrets, err := templates.GetSecrets(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	templateValues := templates.ValuesTemplate{
		Branch:  branch.Name,
		Secrets: secrets,
	}
	err = templates.WriteValues(path, templateValues)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done writing file")
	//

	releaseName := utils.ComposeStrings("todelete-", branch.Name)
	namespace := utils.ComposeStrings("ops-", branch.Name)

	err = helm.InstallOrUpgradeRelease(path, releaseName, namespace)

	fmt.Println("Installing release....")

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
