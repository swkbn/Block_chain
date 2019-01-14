package main

const genesisInfo = "hello！"
//引入区块
//使用数组来模拟区块连
type BlockChain struct {
	Blocks []*Block
}
//创建区块连
/*主要作用就是创建区块连并且，添加创世区块*/
func NewBlockChain()*BlockChain  {

	genesisBlock:=NewBlock(genesisInfo,[]byte{})
	//创建区块连结构，一般会在创建的时候添加一个区块，称之为创世区块
	bc:=BlockChain{
		Blocks:[]*Block{genesisBlock},
	}
	return &bc

}

//添加区块
func (bc *BlockChain)AddBlock(data string)  {
	//创建一个区块，前区块的哈希值从bc的最后一个块元素即可获取
	lasBlock :=bc.Blocks[len(bc.Blocks)-1]//获取最后一个区块
	//即将添加的区块的前哈希值，就是bc中的最后区块的Hahs字段的值
	prevHash:=lasBlock.Hash
	//创建一个新的区块
	newBlock:=NewBlock(data,prevHash)
	//append到区块莲的Blocks数组中
	bc.Blocks=append(bc.Blocks,newBlock)

}
