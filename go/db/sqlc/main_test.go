package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/foyez/go/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err := sql.Open(dbDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	code := m.Run()
	err = testDB.Close()
	if err != nil {
		log.Println("error closing db:", err)
	}
	os.Exit(code)
}
