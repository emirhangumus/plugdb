package main

import (
	"fmt"
	"os"
)

func createTable(state State) (string, error) {
	// create <database> <table>
	if len(state.SplitBuf) != 3 {
		_, err := state.Conn.Write([]byte("Invalid number of arguments\n"))
		if err != nil {
			return "", err
		}
		return "", nil
	}

	tableFile := append(state.SplitBuf[2], ".plugd"...)

	fmt.Println(dbPath + string(state.SplitBuf[1]) + "/" + string(tableFile))

	// create the database if it does not exist
	if _, err := os.Stat(dbPath + string(state.SplitBuf[1])); os.IsNotExist(err) {
		err := os.Mkdir(dbPath+string(state.SplitBuf[1]), 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	// create the table if it does not exist
	if _, err := os.Stat(dbPath + string(state.SplitBuf[1]) + "/" + string(tableFile)); os.IsNotExist(err) {
		_, err := os.Create(dbPath + string(state.SplitBuf[1]) + "/" + string(tableFile))
		if err != nil {
			fmt.Println(err)
		}
	}

	_, err := state.Conn.Write([]byte("Table created\n"))
	if err != nil {
		return "", err
	}

	return string(state.SplitBuf[0]), nil
}
