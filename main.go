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

func (b *blockchain) getLastHash() string {
	return b.blocks[len(b.blocks)-1].hash
}

func (b *blockchain) addBlock(data string){
	newBlock := block{data, "", ""}
	if len(b.blocks) > 0 {
		// hash값이 있어야함
		newBlock.prevHash = b.getLastHash()
	}else {
		newBlock.prevHash = ""
	}
	newBlock.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(data + newBlock.prevHash)))
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks(){
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("PrevHash: %s\n\n", block.prevHash)
	}
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

	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}