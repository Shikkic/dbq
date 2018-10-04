package equality

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func queryListOfTables(db *sql.DB) ([]string, error) {
	sqlQuery := `SELECT *
				FROM
				pg_catalog.pg_tables
				WHERE
				schemaname != 'pg_catalog'
				AND schemaname != 'information_schema';
				`

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dbTables := []string{}
	var tableName string
	var dummyVariable *string
	for rows.Next() {
		err := rows.Scan(
			&dummyVariable,
			&tableName,
			&dummyVariable,
			&dummyVariable,
			&dummyVariable,
			&dummyVariable,
			&dummyVariable,
			&dummyVariable,
		)
		if err != nil {
			return nil, err
		}
		dbTables = append(dbTables, tableName)
	}

	return dbTables, nil
}

func queryTableRowCount(db *sql.DB, tableName string) (int, error) {
	sqlQuery := fmt.Sprintf("SELECT count(*) FROM %s;", tableName)

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var tableRowCount int
	for rows.Next() {
		err := rows.Scan(&tableRowCount)
		if err != nil {
			return 0, err
		}
	}

	return tableRowCount, nil
}

func connectToDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
