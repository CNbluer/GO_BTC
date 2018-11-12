package mode

import (
	"bytes"
	"encoding/gob"
	"crypto/sha256"
	"fmt"
	"os"
)

const reward float64=12.5
type TXInput struct {
	//引用的交易ID
	TXId []byte
	//引用的output索引
	Index int64

	Sign []byte
}

type TXOutput struct {
	Value float64
	PubkeyHash []byte
}

type UTXOinfo struct {
	TXId []byte
	Index int64
	Txoutput *TXOutput
}

type Transaction struct {
	TXId []byte

	TXInputs []*TXInput

	TXOutputs []*TXOutput
}

func (TX *Transaction)SetIdhash()  {
	var buffer bytes.Buffer
	encoder:=gob.NewEncoder(&buffer)
	encoder.Encode(TX)
	hash:=sha256.Sum256(buffer.Bytes())
	TX.TXId=hash[:]
}

func NewcoinBasetx(adress,data []byte)*Transaction  {
	input:=TXInput{nil,-1,data}
	output:=TXOutput{reward,adress}

	txtemp:=Transaction{nil,[]*TXInput{&input},[]*TXOutput{&output}}
	txtemp.SetIdhash()
	return &txtemp
}

//找寻指定地址的所有可用的OUTPUTS
func (bc *BlockChain)FindMyUtxos(address string)[]*UTXOinfo {
	//开启迭代器以遍历全部区块链
	it:=bc.NewblockChainiterator()
	//用来存储已经消耗的即input的索引
	spentutxos:=make(map[string][]int64)
	//用来存储未被花费的
	var TXOutputs []*UTXOinfo
	for  {
		block:=it.Next()
		for _,tx:=range block.Transactions{
			OUTPUT:
			for i,v:=range tx.TXOutputs{
				if string(v.PubkeyHash)==address {
					key:=string(tx.TXId)
					if len(spentutxos[key])!=0 {
						for _,v:=range spentutxos[key]{
							if int64(i)==v {
								continue OUTPUT
							}
						}
					}
					txoutput:=UTXOinfo{tx.TXId,int64(i),v}
					TXOutputs=append(TXOutputs,&txoutput)
				}
			}
			for _,v:=range tx.TXInputs{
				if address==string(v.Sign) {
					key:=string(v.TXId)
					spentutxos[key]=append(spentutxos[key],v.Index)
				}
			}

		}
		if len(block.PrevBlockHash)==0 {
			break
		}
	}
	return TXOutputs
}
//根据可用outputs得到余额
func (bc *BlockChain)GetBalance(address string)float64  {
	outputinfos:=bc.FindMyUtxos(address)
	total:=0.0
	for _,v:=range outputinfos{
		total+=v.Txoutput.Value
	}
	return total
}

//普通转账函数
func (bc *BlockChain)NewTransaction(from,to string,amount float64)*Transaction  {
	needUTXOs,calc:=bc.FindNeedUtxos(from,amount)
	if calc<amount {
		fmt.Println("余额不足，交易创建失败")
		os.Exit(-1)
	}
	zhaoling:=calc-amount
	var tx *Transaction
	var inputs []*TXInput
	for k,arry:=range needUTXOs{
		for _,v:=range arry{
			txput:=TXInput{[]byte(k),v,[]byte(from)}
			inputs=append(inputs,&txput)
		}
	}
	output:=TXOutput{amount,[]byte(to)}
	output1:=TXOutput{zhaoling,[]byte(from)}
	outputs:=make([]*TXOutput,0)
	outputs=append(outputs,&output1,&output)
	tx=&Transaction{nil,inputs,outputs}
	tx.SetIdhash()
	return tx
}

func (bc *BlockChain)FindNeedUtxos(from string,amount float64)(map[string][]int64,float64)  {
	needUTXOS:=make(map[string][]int64)
	Outputsinfo:=bc.FindMyUtxos(from)
	total:=0.0
	for _,v:=range Outputsinfo {
	key:=string(v.TXId)
	total+=v.Txoutput.Value
	needUTXOS[key]=append(needUTXOS[key],v.Index)
		if total>=amount {
			break
		}
	}
	return needUTXOS,total
}



