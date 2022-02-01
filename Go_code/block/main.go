package block

import (
	"fmt"
	""
)

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First data after genesis")
	chain.AddBlock("Second data after genesis")
	chain.AddBlock("Third data after genesis")

	for _, bl := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", bl.PrevHash)
		fmt.Printf("Data in block: %s\n", bl.Data)
		fmt.Printf("Hash: %x\n", bl.Hash)
		fmt.Println()
	}
}