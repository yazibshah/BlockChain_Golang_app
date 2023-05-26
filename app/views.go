package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Post represents a post in the blockchain
type Post struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Index   int    `json:"index"`
	Hash    string `json:"hash"`
}

// Blockchain represents the entire blockchain
type Blockchain struct {
	Chain []Block `json:"chain"`
}

// Block represents a block in the blockchain
type Block struct {
	Index        int64         `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previous_hash"`
}

// Transaction represents a transaction in a block
type Transaction struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

var (
	connectedNodeAddress = "http://127.0.0.1:8000"
	posts                []Post
)

func fetchPosts() {
	getChainAddress := fmt.Sprintf("%s/chain", connectedNodeAddress)
	response, err := http.Get(getChainAddress)
	if err != nil {
		fmt.Println("Error fetching chain:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var chain Blockchain
		if err := json.NewDecoder(response.Body).Decode(&chain); err != nil {
			fmt.Println("Error parsing chain:", err)
			return
		}

		content := []Post{}
		for _, block := range chain.Chain {
			for _, tx := range block.Transactions {
				tx.Index = block.Index
				tx.Hash = block.PreviousHash
				content = append(content, tx)
			}
		}

		posts = sortPostsByTimestamp(content)
	}
}

func sortPostsByTimestamp(posts []Post) []Post {
	// Implementation of sortPostsByTimestamp function
	return posts
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fetchPosts()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":         "YourNet: Decentralized content sharing",
			"posts":         posts,
			"node_address":  connectedNodeAddress,
			"readable_time": timestampToString,
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		var postObject Transaction
		if err := c.ShouldBindJSON(&postObject); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newTxAddress := fmt.Sprintf("%s/new_transaction", connectedNodeAddress)
		_, err := http.Post(newTxAddress, "application/json", nil)
		if err != nil {
			fmt.Println("Error submitting transaction:", err)
			return
		}

		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":8080")
}

func timestampToString(epochTime int64) string {
	return time.Unix(epochTime, 0).Format("15:04")
}
