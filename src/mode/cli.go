package mode

import (
	"os"
	"fmt"
	"strconv"
)

type Cli struct {

}

const Usage = `
	请按照如下格式输入命令行
	createBlockchain address ADDRESS---"create a blockchain if is not exits"
   Send AMOUNT from ADDRESS to ADDRESS miner ADDRESS---"add a transaction in the newblock"
    printChain---"print block Chain"
	getBalance address ADDRESS--"get the number of left money in the address"
`
func (cli *Cli)Run()  {
	if len(os.Args)<2{
		fmt.Println(Usage)
		os.Exit(-1)
	}
	cmd:=os.Args[1]
	switch cmd {
	case "Send":
		if len(os.Args)>8&&os.Args[1]=="Send" {
			amount,_:=strconv.ParseFloat(os.Args[2],64)
			from:=os.Args[4]
			to:=os.Args[6]
			miner:=os.Args[8]
			if from==""||to==""||miner=="" {
				fmt.Println("数据不能为空")
				os.Exit(-1)
			}
			cli.Send(from,to,miner,amount)
		}else {
			fmt.Println(Usage)
		}
	case "printChain":
		cli.Printchain()
	case "createBlockchain":
		if len(os.Args)>3&&os.Args[2]=="address" {
			if os.Args[3]=="" {
				fmt.Println("数据不能为空")
				os.Exit(-1)
			}
			cli.Createchain(os.Args[3])
		}else {
			fmt.Println(Usage)
		}
	case "getBalance":
		if len(os.Args)>3&&os.Args[2]=="address" {
			if os.Args[3]=="" {
				fmt.Println("数据不能为空")
				os.Exit(-1)
			}

			cli.getBalance(os.Args[3])
		}else {
			fmt.Println(Usage)
		}

	default:
		fmt.Println(Usage)
	}
}

func (cli *Cli)Printchain()  {
	bc:=Getblockchainjbk()
	it:=bc.NewblockChainiterator()
	for  {
		block:=it.Next()
		fmt.Println(" ============== =============")
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("MerkleRoot : %x\n", block.MerkelRoot)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Difficuty : %d\n", block.Difficulty)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		//fmt.Printf("Data : %s\n", block.Data)
		if len(block.PrevBlockHash)==0 {
			fmt.Println("that is all")
			break
		}
	}
}

func (cli *Cli)Send(from,to,miner string,amount float64)  {
	bc:=Getblockchainjbk()
	ctx:=bc.NewTransaction(from,to,amount)
	mtx:=NewcoinBasetx([]byte(miner),[]byte{})
	txs:=[]*Transaction{ctx,mtx}
	bc.AddBlock(txs)
	fmt.Println("上链成功")
}

func (cli *Cli)Createchain(address string)  {
	bc:=NewblockChain(address)
	defer bc.db.Close()
}

func (cli *Cli)getBalance(address string)  {
	bc:=Getblockchainjbk()
	fmt.Println(bc.GetBalance(address))
}



