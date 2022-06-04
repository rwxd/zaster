package internal

import "os"

type JSONDatabase struct {
	filePath string
}

func (database *JSONDatabase) checkPathExists() bool {
	info, err := os.Stat(database.filePath)
	if os.IsNotExist(err) {
		return false
	} else if info.IsDir() {
		return false
	}
	return true
}

// func (database *JSONDatabase) createNewDatabaseFile() (err error) {
//
// }
