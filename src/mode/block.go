package mode

import (
	"time"
)

type Block struct {
	Version uint64		//版本号
	MerkelRoot []byte	//这是一个哈希值，后面V5用到
	TimeStamp uint64 	//时间戳
	Difficulty uint64	//难度值
	Nonce uint64		//挖矿所找的随机数
	PrevBlockHash []byte
	Hash []byte
	Transactions []*Transaction
}

func uint2byte(num uint64)[]byte  {
	//var buff bytes.Buffer
	//err:=binary.Write(&buff,binary.BigEndian,&num)
	//if err!=nil {
	//	panic(err)
	//}
	temp:=string(num)
	return []byte(temp)
}

//此函数为第一代版本使用，后被pow代替，只有挖矿之后才能知道真正的哈希值
//func (this *Block)SetHash()  {
//	allinfo:=[][]byte{
//		uint2byte(this.Version),
//		this.MerkelRoot,
//		uint2byte(this.TimeStamp),
//		uint2byte(this.Difficulty),
//		uint2byte(this.Nonce),
//		this.PrevBlockHash,
//		this.Data,
//	}
//	info:=bytes.Join(allinfo,[]byte{})
//	hash:=sha256.Sum256(info)
//	this.Hash=hash[:]
//}

func NewBlock(txs []*Transaction,prevhash []byte)*Block  {
	block:=Block{
		Version:00,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:difficulty,
		PrevBlockHash:prevhash,
		Transactions:txs,
	}
	pow:=NewPOW(block)
	pow.block=block
	nonce,hash:=pow.run()
	block.Nonce=nonce
	block.Hash=hash
	return &block
}




