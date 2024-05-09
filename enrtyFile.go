package main

import (
	"fmt"
	"os"
)

type Database struct {
	Name string
}

type Table struct {
	Name string
}

type DBLook struct {
	Database Database
	Tables   []Table
}

func getSplitedEntryFile(state State, db string) (string, error) {
	// open the file
	file, err := os.Open(dbPath + db + "/entry.plugd")
	if err != nil {
		fmt.Println(err)
	}

	// read the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// read the file
	fileContent := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContent)
	if err != nil {
		fmt.Println(err)
	}

	// generate the database look
	dbLook := DBLook{}
	dbLook.Database.Name = db
	dbLook.Tables = []Table{}

	return string(fileContent), nil
}
