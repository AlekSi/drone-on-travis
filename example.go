package example

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // register driver
)

func GetVersion(db *sql.DB) (string, error) {
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	return version, err
}
