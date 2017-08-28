package example

import (
	"database/sql"
	"testing"
	"time"
)

func getDB(t *testing.T) (db *sql.DB) {
	var err error
	for i := 0; i < 30; i++ {
		db, err = sql.Open("mysql", "root@tcp(mysql:3306)/mysql")
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			return db
		}
		if db != nil {
			db.Close()
		}
		time.Sleep(time.Second)
	}
	t.Fatalf("Failed to connect to db: %s", err)
	panic("not reached")
}

func TestGetVersion(t *testing.T) {
	db := getDB(t)
	version, err := GetVersion(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}
