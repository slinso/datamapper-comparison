package test_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/slinso/datamapper-comparison/gojet"
	"github.com/slinso/datamapper-comparison/testdata"
)

func TestGoJet(t *testing.T) {
	// Create a new mock database and a service that uses it
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	mock.ExpectQuery("SELECT (.+) FROM films").WillReturnRows(sqlmock.NewRows(testdata.Columns).FromCSVString(testdata.Rows))

	// Execute query and store result
	var dest []struct {
		gojet.Actor

		Films []struct {
			gojet.Film

			Language   gojet.Language
			Categories []gojet.Category
		}
	}

	// use qrm and sql-mock together
	_, err = qrm.Query(context.Background(), mockDB, "SELECT * FROM films", nil, &dest)
	if err != nil {
		t.Fatal(err)
	}

	snaps.MatchSnapshot(t, dest)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	var dest2 []struct {
		gojet.Category

		Films  []gojet.Film
		Actors []gojet.Actor
	}

	mock.ExpectQuery("SELECT (.+) FROM films").WillReturnRows(sqlmock.NewRows(testdata.Columns).FromCSVString(testdata.Rows))
	// use qrm and sql-mock together
	_, err = qrm.Query(context.Background(), mockDB, "SELECT * FROM films", nil, &dest2)
	if err != nil {
		t.Fatal(err)
	}

	snaps.MatchSnapshot(t, dest2)
}
