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

type DBConfig struct {
	Host        string `json:"host,omitempty" validate:"required"`
	Port        int    `json:"port,omitempty" validate:"required"`
	User        string `json:"user,omitempty" validate:"required"`
	Password    string `json:"password,omitempty" validate:"required"`
	Dbname      string `json:"dbname,omitempty" validate:"required"`
	RowsPerFile int    `json:"rowsPerFile,omitempty" validate:"required"`
}

var dbConfig DBConfig
var db *sql.DB

func Connect(dbConfigRequest DBConfig) {
	log.Info().Msg("Loading the DB config")
	if dbConfigRequest == (DBConfig{}) {
		log.Info().Msg("Loading the Default DB config")
		dbConfig = DBConfig{Host: host, Port: port, User: user, Password: password, Dbname: dbname}
	} else {
		log.Info().Msg("Loading the Custom DB config")
		dbConfig = dbConfigRequest
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

	// open database
	log.Info().Msg("Trying to open the DB connection")
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		CheckError(err)
	}

	// listRecords(1, db)
	// close database
	// defer db.Close()

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

func ListRecords(page int, tableName string, count int) [][]string {

	limit := dbConfig.RowsPerFile
	offset := limit * (page - 1)
	SQL := `SELECT * FROM ` + tableName + ` LIMIT $2 OFFSET $1`
	rows, err := db.Query(SQL, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil
	}
	cols, _ := rows.Columns()

	data := make([][]string, count+1)

	// rows.Close()

	rowCount := 0
	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)
		row := make([]string, 0)
		if rowCount == 0 {
			data[rowCount] = cols
			rowCount = rowCount + 1
		}
		for i, colName := range cols {
			row = append(row, columns[i])
			log.Info().Msg(colName)
		}
		data[rowCount] = row
		rowCount = rowCount + 1
	}

	return data
}

func GetTableNames() []string {
	var tableName string
	var tables []string
	SQL := "SELECT table_name FROM information_schema.tables	WHERE table_schema = 'public'	ORDER BY table_name;"
	rows, err := db.Query(SQL)

	if err != nil {
		log.Error().Err(err)

	}
	for rows.Next() {
		rows.Scan(&tableName)
		tables = append(tables, tableName)
	}
	return tables
}

func CheckCount(tableName string) (count int) {
	SQL := `SELECT COUNT(*) as count FROM ` + tableName
	rows, err := db.Query(SQL)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&count)
		log.Error().Err(err)
	}
	return count
}
