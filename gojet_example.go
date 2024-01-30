package datamapper

import "github.com/slinso/datamapper-comparison/gojet"

type GoJetExample struct{}

func (GoJetExample) Insert() {
	f := gojet.Film{
		ID:   100,
		URL:  "http://www.postgresqltutorial.com",
		Name: "PostgreSQL Tutorial",
	}

	gojet.Film.InsertStmt().Exec(db)

	f.InsertStmt().Exec(db)

	insertStmt := Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
		MODEL(turorial).
		MODEL(google).
		MODEL(yahoo)
}
