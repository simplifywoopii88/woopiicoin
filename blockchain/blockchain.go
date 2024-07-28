package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data string
	Hash string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks - 1].Hash
}

func createBlock(data string) *block{
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	// 변수 b가 초기화 되었는지 확인
	// 초기화가 안되었다면 초기화
	if b == nil {
		once.Do(func() {
			b = &blockchain{} // blockchain 인스턴스 생성
			b.AddBlock("Genesis")
		})
	}

	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}