# Go Blockchain with Redis and LevelDB

This is a simple implementation of a blockchain in Go, with Redis for inter-client communication and LevelDB for local persistence.

## Features

- **Blockchain:** Implements a basic blockchain structure with blocks containing an index, previous hash, current hash, timestamp, and data.
- **Genesis Block:** Automatically creates a genesis block when the blockchain is first initialized.
- **Persistence with LevelDB:** Persists blockchain data locally using LevelDB.
- **Communication with Redis:** Uses Redis to broadcast new blocks to connected clients.

## Requirements

- Go installed on your machine.
- Redis server running (update connection details in `blockchain.go`).
- [LevelDB Go Package](https://pkg.go.dev/github.com/syndtr/goleveldb/leveldb).

## Getting Started

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd go-blockchain

2. Install dependencies:
    ```bash
    go get -u github.com/go-redis/redis/v8
    go get -u github.com/syndtr/goleveldb/leveldb

3. Run the main program:
    ```bash
    go run main.go

## Configuration
- Update the Redis connection details in blockchain.go if your Redis server is running on a different address.
- Adjust LevelDB settings based on your requirements.

## Contributing
- Contributions are welcome! Please create a new branch for your changes and submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
