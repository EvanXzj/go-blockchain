package main

import (
	"fmt"
	"strconv"
)

func main() {
	blockchain := NewBlockChain()

	blockchain.AddBlock("Send 1 BTC to Ivan")
	blockchain.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range blockchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
