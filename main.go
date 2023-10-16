// main.go
package main

import (
	"fmt"
	"time"

	"github.com/diegocastro-r/go-blockchain/blockchain"
)

func main() {
	// Initialize Redis
	blockchain.InitializeRedis()

	fmt.Println("Redis initialized")

	// Create sender and receiver wallets
	sender, err := blockchain.CreateWallet()
	if err != nil {
		fmt.Println("Error creating sender wallet:", err)
		return
	}

	receiver, err := blockchain.CreateWallet()
	if err != nil {
		fmt.Println("Error creating receiver wallet:", err)
		return
	}

	fmt.Println("Wallets created")
	fmt.Println("Sender: ", blockchain.FormatAddress(sender.PublicKey))
	fmt.Println("Receiver: ", blockchain.FormatAddress(receiver.PublicKey))

	// Initialize blockchain with a custom genesis block
	genesisBlock := blockchain.CreateGenesisBlock()
	customBlockchain := blockchain.NewBlockchainWithGenesis(genesisBlock)
	customBlockchain.PrintBlockchain()
	fmt.Println("Blockchain initialized")

	// Mine new blocks at regular intervals
	for {
		// Add a new block with a transaction from sender to receiver
		customBlockchain.AddBlock("Hello, Blockchain!", sender, receiver, 10)

		fmt.Println("Block mined")

		// Print the blockchain
		customBlockchain.PrintBlockchain()

		// Add a delay before mining the next block
		time.Sleep(10 * time.Second)
	}
}
