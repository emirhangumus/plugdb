package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"emirhangumus.com/plugdb/main/contants"
	"emirhangumus.com/plugdb/main/structs"
)

func createDatabase(state structs.State) (string, error) {
	// create database <name>
	if len(state.SplitBuf) != 2 {
		_, err := state.Conn.Write([]byte("Invalid number of arguments\n"))
		if err != nil {
			return "", err
		}
		return "", nil
	}

	cwd, _ := os.Getwd()
	err := os.Mkdir(contants.DBPath+string(state.SplitBuf[1]), os.ModePerm) // you might want different file access, this suffice for this example
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created %s at %s\n", string(state.SplitBuf[1]), contants.DBPath+string(state.SplitBuf[1]))
	}

	path := filepath.Join(cwd, contants.DBPath+string(state.SplitBuf[1]), contants.EntryFile)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created %s at %s\n", contants.EntryFile, path)
	}

	defer file.Close()

	content := strings.Replace(contants.InitContent, "{0}", string(state.SplitBuf[1]), -1)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}

	return string(state.SplitBuf[1]), nil
}
