package main

import (
	"fmt"
	"strconv"

	"github.com/tensor-programming/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Gensis")
	chain.AddBlock("3rd")

	for _, block := range chain.Blocks {
		fmt.Printf("Previos Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
