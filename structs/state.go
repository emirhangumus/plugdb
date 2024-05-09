package structs

import "net"

type State struct {
	Conn       net.Conn
	Buf        []byte
	TrimmedBuf []byte
	SplitBuf   [][]byte
	History    []byte
}
