package mode

import (
	"bolt"
	"log"
	"encoding/gob"
	"bytes"
	"fmt"
	"os"
)

const genesisInfo = "2009年1月3日，财政大臣正处于实施第二轮银行紧急援助的边缘"

type BlockChain struct {
	//操作数据库的句柄
	db *bolt.DB
	//存储最后一个块的哈希
	lasthash []byte
}

//将Block序列化的函数
func (thisblock *Block)Serialize()[]byte  {
	var buffer bytes.Buffer
	encoder:=gob.NewEncoder(&buffer)
	err:=encoder.Encode(thisblock)
	if err!=nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//反序列化
func Deserilalize(data []byte)*Block  {
	var buffer bytes.Buffer
	var block Block
	buffer.Write(data)
	decoder:=gob.NewDecoder(&buffer)
	decoder.Decode(&block)

	decoder.Decode(block)
	return &block
}

func IfchainExist()bool  {
	_,err:=os.Stat("blockChain.db")
	if os.IsNotExist(err){
		return false
	}else {
		return true
	}

}

//在数据库中创建区块链
func NewblockChain(address string)*BlockChain {
	if IfchainExist(){
		fmt.Println("区块链已经存在，无需创建,可添加或遍历")
		os.Exit(-1)
	}
	var lasthash []byte
	db,err:=bolt.Open("blockChain.db",0600,nil)
	if err !=nil{
		log.Panic(err)
	}
	//defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("blockBucket"))
			bucket,err=tx.CreateBucket([]byte("blockBucket"))
			if err!=nil {
				log.Panic(err)
			}
			transaction:=NewcoinBasetx([]byte(address),[]byte(genesisInfo))
			block:=NewBlock([]*Transaction{transaction},[]byte{})
			bytes:=block.Serialize()
			bucket.Put(block.Hash,bytes)
			bucket.Put([]byte("lasthash"),block.Hash)
			lasthash=block.Hash
		return nil
	})
	return &BlockChain{db,lasthash}

}

//拿到区块链接口
func Getblockchainjbk()*BlockChain {
	if !IfchainExist(){
		fmt.Println("还没有链，请先创建")
		fmt.Println(Usage)
		os.Exit(-1)
	}
	var lasthash []byte
	db,err:=bolt.Open("blockChain.db",0600,nil)
	if err !=nil{
		log.Panic(err)
	}
	//这个接口不能关，！！！！！！！！！！
	//defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("blockBucket"))
		lasthash=bucket.Get([]byte("lasthash"))
		return nil
	})
	return &BlockChain{db,lasthash}
}

func (this *BlockChain)AddBlock(txs []*Transaction)  {
	newblock:=NewBlock(txs,this.lasthash)
	this.db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("blockBucket"))
		if bucket==nil {
			fmt.Println("please new a blockchain firstly")
			log.Panic("no bucket")
		}

		bytes:=newblock.Serialize()
		bucket.Put(newblock.Hash,bytes)
		bucket.Put([]byte("lasthash"),newblock.Hash)
		return nil
	})
	this.lasthash=newblock.Hash
}



