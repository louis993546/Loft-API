package database

import "database/sql"

func InitializeDatabase(db *sql.DB) {
	panic("not implemented")
}

// getSchemaVersion goes to db and get's the schema version from it.
func GetSchemaVersion(db *sql.DB) (int, error) {
	panic("not implemented")
}

func PerformDatabaseMigration(currentVersion int, compatibleVersion int) {
	panic("not implemented")
}
