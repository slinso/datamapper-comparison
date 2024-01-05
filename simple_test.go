package test_test

import "testing"

// write a test case querying the database
func TestDatabase(t *testing.T) {
	stmt := `SELECT film_id FROM dvds.film`

	// execute the query
	rows, err := db.Query(stmt)
	if err != nil {
		t.Fatal(err)
	}

	// iterate over the rows
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(id)
	}
}
