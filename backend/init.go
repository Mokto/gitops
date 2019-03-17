package backend

import (
	"fmt"
	"gitops/backend/config"
	"gitops/backend/controllers"
	"gitops/backend/dao"
	"gitops/backend/services/github"
)

// Init the frontend and api backend
func Init() {
	dao.InitAllTables()

	repos := config.GetRepositories()
	err := dao.InsertManyRepositories(&repos)
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		branches, err := github.GetBranches(repo.Organization, repo.Name)
		if err != nil {
			panic(err)
		}
		err = dao.InsertManyBranches(&branches, repo.ID)
		if err != nil {
			panic(err)
		}
	}

	//for _, branch := range branches {
	//	err = controllers.InstallRepositoryBranch(repository, branch)
	//	if err != nil {
	//		return err
	//	}
	//}
	err = controllers.InstallAllRepositories()
	fmt.Println(err)
}
