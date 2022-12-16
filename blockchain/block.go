package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash          []byte
	Data          []byte
	PrevBlockHash []byte
}

func (b *Block) DerivedHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DerivedHash()
	return block
}

func (chain *Blockchain) Addblock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
