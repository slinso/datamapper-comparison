package test_test

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

var db *sql.DB

func loadData() error {
	// open dump file
	f, err := os.Open("./db/mysql/dvds.sql")
	if err != nil {
		return err
	}

	// read file
	stmt, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// execute the dump
	_, err = db.Exec(string(stmt))
	if err != nil {
		return err
	}

	return nil
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("mysql", "8", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true&multiStatements=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	err = loadData()
	if err != nil {
		log.Fatalf("could not load the data: %s", err)
	}

	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
