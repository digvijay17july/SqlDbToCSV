package api

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "dvdrental"
)

func Connect() {
	log.Info().Msg("Loading the DB config")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	log.Info().Msg("Trying to open the DB connection")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		CheckError(err)
	}

	listRecords(1, db)
	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func listRecords(page int, db *sql.DB) map[string]string {

	limit := 10
	offset := limit * (page - 1)

	SQL := `SELECT "id","nome" FROM "clientes" ORDER BY "id" LIMIT $2 OFFSET $1`

	rows, err := db.Query(SQL, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil
	}
	cols, _ := rows.Columns()

	data := make(map[string]string)

	if rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}
	}
	for key, vlue := range data {
		log.Info().Msg(key + ":" + vlue)
	}
	return data

}
