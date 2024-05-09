package main

import (
	"bytes"
	"fmt"
	"net"
)

type State struct {
	Conn       net.Conn
	Buf        []byte
	TrimmedBuf []byte
	SplitBuf   [][]byte
	History    []byte
}

func main() {
	// Listen on port 5001
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 5001...")

	for {
		// Wait for a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		// Handle the connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// open a command line interface
	fmt.Println("Connection established")
	state := State{Conn: conn, History: []byte("")}

	for {
		// make a buffer to hold incoming data
		state.Buf = make([]byte, 1024)
		// read the incoming connection into the buffer
		conn.Write([]byte("-> "))
		reqLen, err := conn.Read(state.Buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		if reqLen > 0 {

			state.TrimmedBuf = bytes.TrimRight(state.Buf, "\x00")      // Trim trailing null bytes
			state.TrimmedBuf = bytes.TrimRight(state.TrimmedBuf, "\n") // Trim trailing newline

			err = command_execution(state)
			if err != nil {
				fmt.Println("Error executing command:", err.Error())
				break
			}
			state.Buf = nil
		}
	}
}
