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

type blockchain struct {
	blocks []block
}

func (b *blockchain) addBlock(data string){
	newBlock := block{data, "", ""}
	if len(block)


}

func (b *blockchain) listBlocks(){}

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

	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}