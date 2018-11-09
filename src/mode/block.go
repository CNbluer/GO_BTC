package mode

import (
	"bytes"
	"crypto/sha256"
	"time"
	"encoding/binary"
)

type Block struct {
	Version uint64		//版本号
	MerkelRoot []byte	//这是一个哈希值，后面V5用到
	TimeStamp uint64 	//时间戳
	Difficulty uint64	//难度值
	Nonce uint64		//挖矿所找的随机数
	PrevBlockHash []byte
	Data []byte
	Hash []byte
}

func uint2byte(num uint64)[]byte  {
	var buff bytes.Buffer
	err:=binary.Write(&buff,binary.BigEndian,&num)
	if err!=nil {
		panic(err)
	}
	return buff.Bytes()
}

func (this *Block)SetHash()  {
	allinfo:=[][]byte{
		uint2byte(this.Version),
		this.MerkelRoot,
		uint2byte(this.TimeStamp),
		uint2byte(this.Difficulty),
		uint2byte(this.Nonce),
		this.PrevBlockHash,
		this.Data,
	}
	info:=bytes.Join(allinfo,[]byte{})
	hash:=sha256.Sum256(info)
	this.Hash=hash[:]
}

func NewBlock(data string,prevhash []byte)*Block  {
	this:=Block{
		Version:00,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0,
		Nonce:0,
		PrevBlockHash:prevhash,
		Data:[]byte(data),
	}
	this.SetHash()
	return &this
}




