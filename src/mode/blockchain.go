package mode

const genesisInfo = "2009年1月3日，财政大臣正处于实施第二轮银行紧急援助的边缘"

type BlockChain struct {
	Blocks []*Block
}

func NewblockChain()*BlockChain {

	genisisblock:=NewBlock(genesisInfo,[]byte{})
	var this BlockChain
	this.Blocks=append(this.Blocks,genisisblock)
	return &this
}

func (this *BlockChain)AddBlock(data string)  {
	prehash:=this.Blocks[len(this.Blocks)-1].PrevBlockHash
	newblock:=NewBlock(data,prehash)
	this.Blocks=append(this.Blocks,newblock)
}




