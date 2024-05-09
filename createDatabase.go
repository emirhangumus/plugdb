package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const dbPath = "db/"
const entryFile = "entry.plugd"
const initContent string = `[database]
	name={0}
[tables]
`

func createDatabase(state State) (string, error) {
	// create database <name>
	if len(state.SplitBuf) != 2 {
		_, err := state.Conn.Write([]byte("Invalid number of arguments\n"))
		if err != nil {
			return "", err
		}
		return "", nil
	}

	cwd, _ := os.Getwd()
	err := os.Mkdir(dbPath+string(state.SplitBuf[1]), os.ModePerm) // you might want different file access, this suffice for this example
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created %s at %s\n", string(state.SplitBuf[1]), dbPath+string(state.SplitBuf[1]))
	}

	path := filepath.Join(cwd, dbPath+string(state.SplitBuf[1]), entryFile)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created %s at %s\n", entryFile, path)
	}

	defer file.Close()

	content := strings.Replace(initContent, "{0}", string(state.SplitBuf[1]), -1)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}

	return string(state.SplitBuf[1]), nil
}
