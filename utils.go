package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(n int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, n)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
