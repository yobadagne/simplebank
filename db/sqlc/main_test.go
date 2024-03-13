package db

import (
	//"context"
	"database/sql"
	"log"
	"os"

	//"os"
	"testing"
)
const (
	dbDriver= "postgres"
	dbSource = "postgresql://root:yobadagne2nd@localhost:5432/simple_bank?sslmode=disable"
)
var testQueries *Queries
var testDB *sql.DB
func TestMain(m *testing.M) {
	var err error
	testDB,err = sql.Open(dbDriver,dbSource)

	if err != nil {
		log.Fatal("cannot conn to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
