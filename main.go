package main

import (
	"fmt"
	"strconv"

	"github.com/Zajicek1/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.Addblock("First block after Genesis")
	chain.Addblock("Second block after Genesis")
	chain.Addblock("Third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
