// blockchain.go
package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var redisClient *redis.Client

const miningDifficulty = 3

type Block struct {
	Index        int
	PrevHash     string
	Hash         string
	TimeStamp    string
	Data         string
	Transactions []*Transaction
}

type Transaction struct {
	From   *ecdsa.PublicKey
	To     *ecdsa.PublicKey
	Amount int
	// Add more fields as needed
}

type Blockchain struct {
	Blocks []*Block
}

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

// InitializeRedis initializes the Redis client.
func InitializeRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Adjust the address based on your Redis server configuration
		DB:   0,
	})
}

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func CreateWallet() (*Wallet, error) {
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		return nil, err
	}
	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func calculateHash(block *Block) string {
	for {
		hash := fmt.Sprintf("%d%s%s%s", block.Index, block.PrevHash, block.TimeStamp, block.Data)
		if strings.HasPrefix(hash, strings.Repeat("0", miningDifficulty)) {
			return hash
		}
		// If the hash doesn't meet the difficulty, try again with a different timestamp
		block.TimeStamp = time.Now().String()
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{CreateGenesisBlock()},
	}
}

func CreateGenesisBlock() *Block {
	return &Block{
		Index:     0,
		PrevHash:  "0",
		Hash:      "genesis-hash",
		TimeStamp: "genesis-timestamp",
		Data:      "genesis-data",
	}
}

func NewBlockchainWithGenesis(genesis *Block) *Blockchain {
	return &Blockchain{
		Blocks: []*Block{genesis},
	}
}

func (bc *Blockchain) AddBlock(data string, sender *Wallet, receiver *Wallet, amount int) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		PrevHash:     prevBlock.Hash,
		TimeStamp:    time.Now().String(),
		Data:         data,
		Transactions: []*Transaction{{From: sender.PublicKey, To: receiver.PublicKey, Amount: amount}},
	}

	// Calculate hash and set it
	newBlock.Hash = calculateHash(newBlock)

	// Append block to the chain
	bc.Blocks = append(bc.Blocks, newBlock)

	// Broadcast the new block to Redis channel
	broadcastBlockToRedis(newBlock)
}

func broadcastBlockToRedis(block *Block) error {
	// Encode block data to JSON
	blockData, err := json.Marshal(block)
	if err != nil {
		return err
	}

	// Publish block data to the Redis channel
	err = redisClient.Publish(context.Background(), "blockchain_channel", string(blockData)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d, PrevHash: %s, Hash: %s, TimeStamp: %s, Data: %s\n",
			block.Index, block.PrevHash, block.Hash, block.TimeStamp, block.Data)
	}
}

func FormatAddress(publicKey *ecdsa.PublicKey) string {
	// Convert the public key to a hexadecimal string
	pubKeyHex := hex.EncodeToString(elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y))

	// Truncate or pad the string to 42 characters
	if len(pubKeyHex) > 42 {
		pubKeyHex = pubKeyHex[:42]
	} else {
		pubKeyHex = fmt.Sprintf("%-42s", pubKeyHex)
	}

	return pubKeyHex
}
