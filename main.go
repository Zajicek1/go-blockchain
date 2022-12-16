package main

import (
	"fmt"

	"blockchain/framework"
)

func main() {
	chain := framework.InitBlockchain()

	chain.Addblock("First block after Genesis")
	chain.Addblock("Second block after Genesis")
	chain.Addblock("Third block after Genesis")

	for _, block := range chain.Blocks {
		// fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
