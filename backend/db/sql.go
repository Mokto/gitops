package db

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func GetConnection() (*sqlite3.Conn, error) {
	return sqlite3.Open("gitops.db")
}