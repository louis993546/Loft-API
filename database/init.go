package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

// InitializeDatabase initialize the db according to init.sql
func InitializeDatabase(db *sql.DB) error {
	absPath, _ := filepath.Abs("../database/init.sql")
	fileByteArray, fileErr := ioutil.ReadFile(absPath)
	if fileErr != nil {
		return fileErr
	}

	_, execErr := db.Exec(fmt.Sprintf("%s", fileByteArray))
	return execErr
}

// GetSchemaVersion goes to db and get's the schema version from it.
func GetSchemaVersion(db *sql.DB) (int, error) {
	rows, queryErr := db.Query("SELECT loft.meta.value FROM loft WHERE loft.meta.key='SCHEMA_VERSION'")
	if queryErr != nil {
		switch queryErr.Error() {
		case "pq: relation \"loft\" does not exist":
			return -1, NewNotFoundError("Schema 'loft' does not exist")
		default:
			//TODO: find out what other kinds of error can happen
			return -1, queryErr
		}
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
		return -1, NewNotFoundError("Database (and meta table) exists, but there is no schema version record")
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
