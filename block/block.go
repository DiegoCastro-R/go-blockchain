package block

import (
	"crypto/ecdsa"
	"fmt"
)

type Transaction struct {
	From   *ecdsa.PublicKey
	To     *ecdsa.PublicKey
	Amount int
	// Add more fields as needed
}

// Add a Transactions field to your Block struct
type Block struct {
	Index        int
	PrevHash     string
	Hash         string
	TimeStamp    string
	Data         string
	Transactions []*Transaction
}

// createGenesisBlock creates the genesis block.package block

func CreateGenesisBlock() *Block {
	return &Block{
		Index:     0,
		PrevHash:  "0",
		Hash:      "genesis-hash",
		TimeStamp: "genesis-timestamp",
		Data:      "genesis-data",
	}
}

// calculateHash calculates the hash of a block.
// Replace with a real hash function.
func CalculateHash(block *Block) string {
	return fmt.Sprintf("%d%s%s%s", block.Index, block.PrevHash, block.TimeStamp, block.Data)
}

// generateKeyPair generates a new key pair.
