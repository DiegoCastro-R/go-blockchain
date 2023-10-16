package main

import (
	"github.com/diegocastro-r/go-blockchain/blockchain"
)

func main() {
	// Initialize blockchain
	blockchain := blockchain.NewBlockchain()

	// Add a new block
	blockchain.AddBlock("Hello, Blockchain!")

	// Print the blockchain
	blockchain.PrintBlockchain()
}
