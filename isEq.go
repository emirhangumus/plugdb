package main

import "bytes"

func isEq(a, b []byte) bool {
	return bytes.Equal(a, b)
}
