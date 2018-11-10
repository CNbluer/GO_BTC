package mode

import (
	"bolt"
	"fmt"
	"os"
)

type BlockchainIterator struct {
	db *bolt.DB
	current_point []byte
}

func (bc *BlockChain)NewblockChainiterator()*BlockchainIterator  {
	var it BlockchainIterator
	it.db=bc.db
	it.current_point=bc.lasthash
	return &it
}
func (it *BlockchainIterator)Next()*Block  {
	var block *Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("blockBucket"))
	if bucket==nil {
		fmt.Println("bucket is nill")
		os.Exit(-1)
	}
	bytes:=bucket.Get(it.current_point)
	block=Deserilalize(bytes)
	return nil
})
	it.current_point=block.PrevBlockHash
	return block
}

