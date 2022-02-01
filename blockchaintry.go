package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

// now we have created a block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

//we are now making method for deriving the hash of the block
func (b *Block) DeriveHash() {
	// joins 2 bytes together with the given separator, in this case it is empty
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// sha256 is simpler than real hash algo
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//now we will create the blocks as per the values given
//9:57
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Fist genesis block", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First data after genesis")
	chain.AddBlock("Second data after genesis")
	chain.AddBlock("Third data after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
