package main

import (
	"bytes"
	"fmt"

	"emirhangumus.com/plugdb/main/structs"
)

func CommandExecution(state structs.State) error {
	if isEq(state.TrimmedBuf, []byte("exit")) {
		fmt.Println("Connection closed")
		state.Conn.Close()
		return nil
	}

	state.SplitBuf = bytes.Split(state.TrimmedBuf, []byte(" "))

	switch {
	case isEq(state.SplitBuf[0], []byte("database")):
		_, err := createDatabase(state)
		if err != nil {
			return err
		}
	case isEq(state.SplitBuf[0], []byte("show")):
		_, err := showDatabases(state)
		if err != nil {
			return err
		}
	case isEq(state.SplitBuf[0], []byte("table")):
		_, err := createTable(state)
		if err != nil {
			return err
		}
	default:
		_, err := state.Conn.Write([]byte("Unknown command\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
