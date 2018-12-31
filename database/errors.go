package database

import (
	"fmt"
	"strings"
)

// ErrorNotFound means none of the tables can be found
type ErrorNotFound struct {
	message string
}

// NewNotFoundError is essentially a constructor for DatabaseNotFoundError
func NewNotFoundError(message string) *ErrorNotFound {
	return &ErrorNotFound{
		message: message,
	}
}

func (e *ErrorNotFound) Error() string {
	return e.message
}

// ErrorCorrupted represents database that said they are vX, but when you check each table, they don't match up
type ErrorCorrupted struct {
	version         int
	missingTables   []string
	incorrectTables []string
}

// NewCorruptedError is essentially a constructor for DatabaseCorruptedError
func NewCorruptedError(schemaVersion int, missingTables []string, incorrectTables []string) *ErrorCorrupted {
	return &ErrorCorrupted{
		version:         schemaVersion,
		missingTables:   missingTables,
		incorrectTables: incorrectTables,
	}
}

func (e *ErrorCorrupted) Error() string {
	return fmt.Sprintf("Database said it's v%d, but the schema doesn't match up:\nMissing table: %s\nIncorrect tables: %s\n", e.version, strings.Join(e.missingTables, ", "), strings.Join(e.incorrectTables, ", "))
}
