package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data string	`json:"data"`
	Hash string	`json:"hash"`
	PrevHash string	`json:"prev_hash,omitempty"`
	Height int	`json:"height"`
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
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

func createBlock(data string) *Block{
	newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
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

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) (*Block, error){
	if height > len(b.blocks){
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}

var ErrNotFound = errors.New("block not found")