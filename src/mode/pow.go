package mode

import (
	"math/big"
	"bytes"
	"crypto/sha256"
)

const difficulty =18

type ProofOfWork struct {
	//工作量证明所需目标值
	target *big.Int
	//区块数据
	block Block
}


func NewPOW(block Block)*ProofOfWork  {
	var this ProofOfWork
	this.block=block
	targetInt:=big.NewInt(1)
	targetInt.Lsh(targetInt,256-difficulty)
	this.target=targetInt
	return &this
}

//添加挖矿算法前需要一个辅助得到哈希值的函数
func (pow *ProofOfWork)PrepareData(nonce uint64)[32]byte  {
	this:=pow.block
	this.HashTransactions()
	allinfo:=[][]byte{
		uint2byte(this.Version),
		this.MerkelRoot,
		uint2byte(this.TimeStamp),
		uint2byte(this.Difficulty),
		uint2byte(nonce),
		this.PrevBlockHash,
		//this.data
		//在真正的比特币中计算的是哈希头，根据梅克尔数
	}
	info:=bytes.Join(allinfo,[]byte{})
	return sha256.Sum256(info)

}

//给到一个梅克尔数计算方法
func (this *Block)HashTransactions()  {
	var temp [][]byte
	for _,v:=range this.Transactions  {
	temp=append(temp,v.TXId)
	}
	hash:=sha256.Sum256(bytes.Join(temp,[]byte{}))
	this.MerkelRoot=hash[:]
}

//添加一个挖矿算法
func (this *ProofOfWork)run()(uint64,[]byte)  {
	var nonce uint64
	var hash [32]byte
	for  {
		hash=this.PrepareData(nonce)
		var hashBigInt big.Int
		hashBigInt.SetBytes(hash[:])

		if hashBigInt.Cmp(this.target)==-1{
			break
		}else {

			nonce++
		}
	}

	return nonce,hash[:]
}

