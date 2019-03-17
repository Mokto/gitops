package dao

import "fmt"

func InitAllTables() {
	err := initRepositoryTable()
	if err != nil {
		fmt.Println(err)
	}
	err = initBranchTable()
	if err != nil {
		fmt.Println(err)
	}
}
