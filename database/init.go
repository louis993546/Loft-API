package database

import (
	"database/sql"
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

	switch len(versions) {
	case 0:
		return -1, NewNotFoundError("")
	case 1:
		return strconv.Atoi(versions[0])
	default:
		// TODO: might want to break down how many types of "corrupt" error there can be
		return -1, NewCorruptedError(-1, []string{}, []string{})
	}
}

func PerformDatabaseMigration(currentVersion int, compatibleVersion int) {
	panic("not implemented")
}
