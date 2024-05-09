package main

import (
	"bytes"
	"fmt"
)

// func printAsciiCodes(input []byte) {
// 	for _, b := range input {
// 		fmt.Printf("%d ", b)
// 	}
// 	fmt.Println()
// }

func command_execution(state State) error {

	if isEq(state.TrimmedBuf, []byte("exit")) {
		fmt.Println("Connection closed")
		state.Conn.Close()
		return nil
	}

	state.SplitBuf = bytes.Split(state.TrimmedBuf, []byte(" "))

	switch {
	// is equal to create
	case isEq(state.SplitBuf[0], []byte("create")):
		_, err := createDatabase(state)
		if err != nil {
			return err
		}
	case isEq(state.SplitBuf[0], []byte("show")):
		_, err := showDatabases(state)
		if err != nil {
			return err
		}
	case isEq(state.SplitBuf[0], []byte("touch")):
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
