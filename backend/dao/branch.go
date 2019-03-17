package dao

import (
	"gitops/backend/db"
	"gitops/backend/models"
)

func initBranchTable() error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = conn.Exec(`DROP TABLE IF EXISTS branches`)
	if err != nil {
		return err
	}
	err = conn.Exec(`CREATE TABLE branches(
		repositoryId TEXT,
		name TEXT,
		FOREIGN KEY (repositoryId)
		REFERENCES repositories(id)
	)`)
	if err != nil {
		return err
	}
	return nil
}

func InsertManyBranches(branches *[]models.Branch, repositoryId string) (err error) {
	for i := range *branches {
		err := InsertOneBranch(&(*branches)[i], repositoryId)
		if err != nil {
			return err
		}
	}
	return
}

func InsertOneBranch(branch *models.Branch, repositoryId string) (err error) {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	branch.RepositoryId = repositoryId
	err = conn.Exec(`INSERT INTO branches(name, repositoryId) VALUES (?, ?)`, branch.Name, repositoryId)
	if err != nil {
		return err
	}
	return nil
}

func FindAllBranchesByRepository(repositoryId string) (branches []models.Branch, err error) {
	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	// Prepare can prepare a statement and optionally also bind arguments
	stmt, err := conn.Prepare(`SELECT name, repositoryId FROM branches WHERE repositoryId = ?`, repositoryId)
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

		branch := models.Branch{}
		err = stmt.Scan(&branch.Name, &branch.RepositoryId)
		if err != nil {
			return nil, err
		}
		branches = append(branches, branch)
	}

	return branches, nil
}
