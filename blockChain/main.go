package main

import (
	"fmt"
)
//简单版本
//6：重构代码
func main() {
	bc:=NewBlockChain()
	//创建区块
	bc.AddBlock("你好，航头")
	bc.AddBlock("再见，航头")
	for i,block:=range bc.Blocks{
		fmt.Printf("==========区块高度：%d ========\n",i)
		fmt.Printf("前区块哈希值：%x\n",block.PrevBlockHash)
		fmt.Printf("当前区块哈希值：%x\n",block.Hash)
		fmt.Printf("区块数据：%x\n",block.Data)
	}
}

