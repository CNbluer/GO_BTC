package main

import (
	"fmt"
	"mode"
	"encoding/hex"
)

func main() {
	blockchain:=mode.NewblockChain()
	blockchain.AddBlock("heheda")
	blockchain.AddBlock("kawayi")
	for i,v:=range blockchain.Blocks{
		fmt.Printf("=================区块高度%d================",i)
		fmt.Println(string(v.PrevBlockHash))
		fmt.Println(hex.EncodeToString(v.Hash))
		fmt.Println(string(v.Data))
	}


}