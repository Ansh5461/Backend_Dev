package main

import (
	"bytes"
	"crypto/sha256"
)

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

func main() {

}
