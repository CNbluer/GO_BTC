package mode

import (
	"os"
	"fmt"
)

type Cli struct {

}

const Usage = `
	请按照如下格式输入命令行
	createBlockchain---"create a blockchain if is not exits"
    addBlock data DATA---"add a block"
    printChain---"print block Chain"
`
func (cli *Cli)Run()  {
	if len(os.Args)<2{
		fmt.Println(Usage)
	}
	cmd:=os.Args[1]
	switch cmd {
	case "addBlock":
		if len(os.Args)>3&&os.Args[2]=="data" {
			data:=os.Args[3]
			if data=="" {
				fmt.Println("数据不能为空")
				os.Exit(-1)
			}
			cli.Addblock(data)
		}else {
			fmt.Println(Usage)
		}
	case "printChain":
		cli.Printchain()
	case "createBlockchain":
		cli.Createchain()
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
		fmt.Printf("Data : %s\n", block.Data)
		if len(block.PrevBlockHash)==0 {
			fmt.Println("that is all")
			break
		}
	}
}

func (cli *Cli)Addblock(data string)  {
	bc:=Getblockchainjbk()
	bc.AddBlock(data)
	fmt.Println("上链成功")
}

func (cli *Cli)Createchain()  {
	bc:=NewblockChain()
	defer bc.db.Close()
}



