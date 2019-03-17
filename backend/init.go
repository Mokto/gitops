package backend

import (
	"fmt"
	"gitops/backend/config"
	"gitops/backend/controllers"
	"gitops/backend/dao"
)

// Init the frontend and api backend
func Init() {
	dao.InitAllTables()

	repos := config.GetRepositories()
	err := dao.InsertManyRepositories(repos)
	if err != nil {
		panic(err)
	}

	err = controllers.InstallAllRepositories()
	fmt.Println(err)
}
