package main

// Blockchain struct
type Blockchain struct {
	blocks []*Block
}

// AddBlock ...
func (bc *Blockchain) AddBlock(data string) {
	prevBlockHash := bc.blocks[len(bc.blocks)-1].Hash
	newBlock := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockChain ...
func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
