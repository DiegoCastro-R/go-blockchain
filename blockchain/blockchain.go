package blockchain

import (
	"fmt"

	"github.com/diegocastro-r/go-blockchain/block"
)

// Blockchain represents the blockchain.
type Blockchain struct {
	Blocks []*block.Block
}

// NewBlockchain initializes a new blockchain.
func NewBlockchain() *Blockchain {
	// Check if blockchain is already initialized
	existingBlocks := loadBlocksFromDatabase()
	if len(existingBlocks) > 0 {
		return &Blockchain{
			Blocks: existingBlocks,
		}
	}

	// If blockchain is empty, create genesis block
	return &Blockchain{
		Blocks: []*block.Block{block.CreateGenesisBlock()},
	}
}

// AddBlock adds a new block to the blockchain.
// loadBlocksFromDatabase fetches blocks from LevelDB.
func loadBlocksFromDatabase() []*block.Block {
	// Implement LevelDB integration here
	// Return existing blocks or an empty slice if not found
	return []*block.Block{}
}

// AddBlock adds a new block to the blockchain and persists it in LevelDB.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &block.Block{
		Index:     prevBlock.Index + 1,
		PrevHash:  prevBlock.Hash,
		TimeStamp: "timestamp", // Replace with actual timestamp logic
		Data:      data,
	}

	// Calculate hash and set it
	newBlock.Hash = block.CalculateHash(newBlock)

	// Append block to the chain
	bc.Blocks = append(bc.Blocks, newBlock)

	// Persist the new block in LevelDB
	persistBlockInDatabase(newBlock)

	// Broadcast the new block to Redis channel
	broadcastBlockToRedis(newBlock)

}

// persistBlockInDatabase stores a block in LevelDB.
func persistBlockInDatabase(block *block.Block) {
	// Implement LevelDB integration here
}

// broadcastBlockToRedis publishes a block to a Redis channel.
func broadcastBlockToRedis(block *block.Block) {
	// Implement Redis integration here
}

// PrintBlockchain prints the blocks in the blockchain.
func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d, PrevHash: %s, Hash: %s, TimeStamp: %s, Data: %s\n",
			block.Index, block.PrevHash, block.Hash, block.TimeStamp, block.Data)
	}
}
