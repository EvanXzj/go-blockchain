package main

import "testing"

func TestNewGenesisBlock(t *testing.T) {
	coinbase := NewCoinbaseTX("xyz", "Genesis Block")

	block := NewGenesisBlock(coinbase)

	if string(block.Transactions[0].Vin[0].ScriptSig) != "Genesis Block" {
		t.Error("TestNewGenesisBlock failed")
	}
}
