package main

import (
	"fmt"
	"os"

	"emirhangumus.com/plugdb/main/contants"
	"emirhangumus.com/plugdb/main/structs"
)

func showDatabases(state structs.State) (string, error) {
	// create database <name>
	if len(state.SplitBuf) == 1 {
		files, err := os.ReadDir(contants.DBPath)
		if err != nil {
			fmt.Println(err)
		}

		for index, file := range files {
			state.Conn.Write([]byte(fmt.Sprintf("%d. ", index+1)))
			state.Conn.Write([]byte(file.Name() + "\n"))
		}
	} else if len(state.SplitBuf) == 2 {
		files, err := os.ReadDir(contants.DBPath + string(state.SplitBuf[1]))
		if err != nil {
			fmt.Println(err)
		}

		for index, file := range files {
			state.Conn.Write([]byte(fmt.Sprintf("%d. ", index+1)))
			state.Conn.Write([]byte(file.Name() + "\n"))
		}
	} else if len(state.SplitBuf) == 3 {
		// the command is show first 1 -> "show" is the command, "first" is the database name, "1" is the number of file the content of the file will be shown
		files, err := os.ReadDir(contants.DBPath + string(state.SplitBuf[1]))
		if err != nil {
			fmt.Println(err)
		}

		file, err := os.ReadFile(contants.DBPath + string(state.SplitBuf[1]) + "/" + files[0].Name())
		if err != nil {
			fmt.Println(err)
		}

		state.Conn.Write(file)
	} else {
		_, err := state.Conn.Write([]byte("Invalid number of arguments\n"))
		if err != nil {
			return "", err
		}
		return "", nil
	}

	return string(state.SplitBuf[0]), nil
}
