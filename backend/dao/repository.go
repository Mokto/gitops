package dao

import (
	"gitops/backend/db"
	"gitops/backend/models"
)

func initRepositoryTable() (error) {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = conn.Exec(`DROP TABLE IF EXISTS repositories`)
	if err != nil {
		return err
	}
	err = conn.Exec(`CREATE TABLE repositories(type TEXT, organization TEXT, name TEXT, fullName TEXT)`)
	if err != nil {
		return err
	}
	return nil
}

func InsertManyRepositories(repositories []models.Repository) (err error) {
	for _, repository := range repositories {
		err := InsertOneRepository(repository)
		if err != nil {
			return err
		}
	}
	return
}

func InsertOneRepository(repo models.Repository) (err error) {
	conn, err := db.GetConnection()
	defer conn.Close()
	if err != nil {
		return err
	}
	err = conn.Exec(`INSERT INTO repositories VALUES (?, ?, ?, ?)`, repo.Type, repo.Organization, repo.Name, repo.FullName)
	if err != nil {
		return err
	}
	return nil
}


func FindAllRepositories() (repositories []models.Repository, err error) {
	conn, err := db.GetConnection()
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	// Prepare can prepare a statement and optionally also bind arguments
	stmt, err := conn.Prepare(`SELECT type, organization, name, fullName FROM repositories`)
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
		err = stmt.Scan(&repository.Type, &repository.Organization, &repository.Name, &repository.FullName)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, repository)
	}

	return repositories, nil
}