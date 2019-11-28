package main

import "testing"

func TestNewGenesisBlock(t *testing.T) {
	block := NewGenesisBlock()

	if string(block.Data) != "Genesis Block" {
		t.Error("TestNewGenesisBlock failed")
	}
}
