package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func InitializeDatabase(db *sql.DB) {
	panic("not implemented")
}

// GetSchemaVersion goes to db and get's the schema version from it.
func GetSchemaVersion(db *sql.DB) (int, error) {
	rows, queryErr := db.Query("SELECT loft.meta.value FROM loft WHERE loft.meta.key='SCHEMA_VERSION'")
	if queryErr != nil {
		return -1, queryErr
	}
	defer rows.Close()

	var versions []string
	for rows.Next() {
		var version string
		if rowErr := rows.Scan(&version); rowErr != nil {
			return -1, rowErr
		}
		versions = append(versions, version)
	}

	if len(versions) != 1 {
		return -1, errors.New("There are more then 1 row with key 'SCHEMA_VERSION', something went wrong")
	}

	return strconv.Atoi(versions[0])
}

func PerformDatabaseMigration(currentVersion int, compatibleVersion int) {
	panic("not implemented")
}
