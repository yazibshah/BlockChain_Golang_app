package main
import "github.com/gin-gonic/gin"

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Block struct {
	Index        int
	Transactions []Transaction
	Timestamp    int64
	PreviousHash string
	Nonce        int
	Hash         string
}

type Transaction struct {
	Author    string
	Content   string
	Timestamp int64
}

type Blockchain struct {
	Difficulty               int
	UnconfirmedTransactions  []Transaction
	Chain                    []Block
	chainMutex               sync.Mutex
	unconfirmedTxMutex       sync.Mutex
}

var (
	blockchain  Blockchain
	peers       = make(map[string]bool)
	chainFile   = "chain.json"
)

func main() {
	loadChainFromDisk()

	http.HandleFunc("/new_transaction", newTransactionHandler)
	http.HandleFunc("/chain", getChainHandler)
	http.HandleFunc("/mine", mineHandler)
	http.HandleFunc("/register_node", registerNodeHandler)
	http.HandleFunc("/register_with", registerWithExistingNodeHandler)
	http.HandleFunc("/add_block", addBlockHandler)
	http.HandleFunc("/pending_tx", getPendingTransactionsHandler)

	go listenForShutdownSignal()

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func loadChainFromDisk() {
	file, err := os.Open(chainFile)
	if err != nil {
		genesisBlock := Block{
			Index:        0,
			Transactions: []Transaction{},
			Timestamp:    0,
			PreviousHash: "0",
			Nonce:        0,
		}
		genesisBlock.Hash = computeHash(genesisBlock)
		blockchain = Blockchain{
			Difficulty:              2,
			UnconfirmedTransactions: []Transaction{},
			Chain:                   []Block{genesisBlock},
		}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&blockchain); err != nil {
		log.Fatalf("Failed to decode blockchain from disk: %s", err)
	}
}

func saveChainToDisk() {
	file, err := os.Create(chainFile)
	if err != nil {
		log.Fatalf("Failed to create chain file: %s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(blockchain); err != nil {
		log.Fatalf("Failed to encode blockchain to disk: %s", err)
	}
}

func computeHash(block Block) string {
	blockJSON, err := json.Marshal(block)
	if err != nil {
		log.Fatalf("Failed to marshal block to JSON: %s", err)
	}
	hash := sha256.Sum256(blockJSON)
	return string(hash[:])
}

// Handlers

func newTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var tx Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	tx.Timestamp = time.Now().Unix()

	blockchain.unconfirmedTxMutex.Lock()
	blockchain.UnconfirmedTransactions = append(blockchain.UnconfirmedTransactions, tx)
	blockchain.unconfirmedTxMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func getChainHandler(w http.ResponseWriter, r *http.Request) {
	blockchain.chainMutex.Lock()
	defer blockchain.chainMutex.Unlock()

	chainData := make([]Block, len(blockchain.Chain))
	copy
