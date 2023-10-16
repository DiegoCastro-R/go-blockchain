package block

import "fmt"

// Block represents a block in the blockchain.
type Block struct {
	Index     int
	PrevHash  string
	Hash      string
	TimeStamp string
	Data      string
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
