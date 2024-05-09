package main

import (
	"fmt"
	"os"

	"emirhangumus.com/plugdb/main/contants"
	"emirhangumus.com/plugdb/main/dblook"
	"emirhangumus.com/plugdb/main/structs"
)

func createTable(state structs.State) (string, error) {
	// create <database> <table>
	if len(state.SplitBuf) != 3 {
		_, err := state.Conn.Write([]byte("Invalid number of arguments\n"))
		if err != nil {
			return "", err
		}
		return "", nil
	}

	tableFile := append(state.SplitBuf[2], ".plugd"...)

	fmt.Println(contants.DBPath + string(state.SplitBuf[1]) + "/" + string(tableFile))

	// create the database if it does not exist
	if _, err := os.Stat(contants.DBPath + string(state.SplitBuf[1])); os.IsNotExist(err) {
		err := os.Mkdir(contants.DBPath+string(state.SplitBuf[1]), 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	// create the table if it does not exist
	if _, err := os.Stat(contants.DBPath + string(state.SplitBuf[1]) + "/" + string(tableFile)); os.IsNotExist(err) {
		_, err := os.Create(contants.DBPath + string(state.SplitBuf[1]) + "/" + string(tableFile))
		if err != nil {
			fmt.Println(err)
		}
	}

	_, err := state.Conn.Write([]byte("Table created\n"))
	if err != nil {
		return "", err
	}

	look, err := dblook.GetDBLook(string(state.SplitBuf[1]))
	if err != nil {
		return "", err
	}

	look.Tables = append(look.Tables, dblook.Table{Name: string(state.SplitBuf[2])})

	// write the look back to the file
	err = dblook.WriteDBLook(look, string(state.SplitBuf[1]))
	if err != nil {
		return "", err
	}

	return string(state.SplitBuf[0]), nil
}
