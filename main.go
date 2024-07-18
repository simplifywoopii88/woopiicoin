package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{
		data:     "Genesis Block",
		hash:     "",
		prevHash: "",
	}
	// genesisBlock.hash = sha256(genesisBlock.data + genesisBlock.prevHash)
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}