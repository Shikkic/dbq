package equality

import (
	"database/sql"
	"errors"
	"log"
)

var blacklistedTables = map[string]bool{
	"alembic": true,
}

// CheckDBSubsetEquality todo
func CheckDBSubsetEquality(sourceDBUrl string, targetDBUrl string) error {
	// Establish connection to source DB.
	sourceDB, err := connectToDB(sourceDBUrl)
	if err != nil {
		return err
	}

	// Establish connection to source DB.
	targetDB, err := connectToDB(targetDBUrl)
	if err != nil {
		return err
	}

	sourceTableNames, err := queryListOfTables(sourceDB)
	if err != nil {
		return err
	}

	targetTableNames, err := queryListOfTables(targetDB)
	if err != nil {
		return err
	}

	ok := checkSourceTablesExistInTargetDB(sourceTableNames, targetTableNames)
	if !ok {
		return errors.New("Mismatch in tables")
	}

	equals, err := checkTableRowCountsMatch(sourceDB, targetDB, sourceTableNames)
	if err != nil {
		return err
	}

	if !equals {
		return errors.New("Row counts did not match up")
	}

	return nil
}

func checkTableRowCountsMatch(sourceDB *sql.DB, targetDB *sql.DB, tableNames []string) (bool, error) {
	var rowCountMismatch bool
	for _, tableName := range tableNames {
		_, ok := blacklistedTables[tableName]
		if ok {
			continue
		}

		sourceRowCount, err := queryTableRowCount(sourceDB, tableName)
		if err != nil {
			return false, err
		}

		targetRowCount, err := queryTableRowCount(targetDB, tableName)
		if err != nil {
			return false, err
		}

		if sourceRowCount != targetRowCount {
			rowCountMismatch = true
			log.Printf("Mismatch row count for table \"%s\"\n\t\tSource db=%d, Target db=%d, Difference=%d\n\n", tableName, sourceRowCount, targetRowCount, (targetRowCount - sourceRowCount))
		}
	}

	if rowCountMismatch {
		return false, nil
	}

	return true, nil
}

func checkSourceTablesExistInTargetDB(sourceTableNames []string, targetTableNames []string) bool {
	sourceTableMap := make(map[string]bool)
	for _, sourceTableName := range sourceTableNames {
		sourceTableMap[sourceTableName] = true
	}

	targetTableMap := make(map[string]bool)
	for _, targetTableName := range targetTableNames {
		targetTableMap[targetTableName] = true
	}

	var missingTable bool
	for sourceTableKey := range sourceTableMap {
		_, ok := targetTableMap[sourceTableKey]
		if !ok {
			log.Printf("Target db missing table: %s\n", sourceTableKey)
			missingTable = true
		}
	}

	if missingTable {
		return false
	}

	return true
}
