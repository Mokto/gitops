package dao

import (
	"fmt"
	"gitops/backend/db"
	"gitops/backend/models"
	"gitops/backend/utils"
)

func initRepositoryTable() error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = conn.Exec(`DROP TABLE IF EXISTS repositories`)
	if err != nil {
		return err
	}
	err = conn.Exec(`CREATE TABLE repositories(
		id TEXT PRIMARY KEY,
		type TEXT,
		organization TEXT,
		name TEXT,
		fullName TEXT
	)`)
	if err != nil {
		return err
	}
	return nil
}

func InsertManyRepositories(repositories *[]models.Repository) (err error) {
	for i := range *repositories {
		err := InsertOneRepository(&(*repositories)[i])
		if err != nil {
			return err
		}
	}
	return
}

func InsertOneRepository(repo *models.Repository) (err error) {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo.ID = utils.ComposeStrings(repo.Type, "/", repo.FullName)
	err = conn.Exec(`INSERT INTO repositories(id, type, organization, name, fullName) VALUES (?, ?, ?, ?, ?)`, repo.ID, repo.Type, repo.Organization, repo.Name, repo.FullName)
	if err != nil {
		return err
	}
	return nil
}

func FindAllRepositories() (repositories []models.Repository, err error) {
	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare(`SELECT id, type, organization, name, fullName FROM repositories`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for {
		hasRow, err := stmt.Step()
		if err != nil {
			return nil, err
		}
		if !hasRow {
			// The query is finished
			break
		}

		repository := models.Repository{}
		err = stmt.Scan(&repository.ID, &repository.Type, &repository.Organization, &repository.Name, &repository.FullName)
		fmt.Println(repository)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, repository)
	}

	return repositories, nil
}
