package main

import (
	"os"
	"fmt"
)

//定义一个CLI结构
//不需要字段
//需要一个方法Run
//例如：addBlock命令
const Usage  = `
		./block createBC "创建区块连数据库"
		./block addBlock "DATA" 创建区块连数据库
		./block printChain "打印区块连数据库"
`

type CLI struct {
	//不需要参数
}

func (cli*CLI)Run()  {
	cmds:=os.Args

	if len(cmds)<2 {
		fmt.Println(Usage)
		os.Exit(1)
	}
	switch cmds[1] {
	case "createBC":
		fmt.Println("创建区块链命令被调用")
		cli.createBc()
	case "addBlock":
		fmt.Println("添加区块命令被调用")
		if len(cmds)!=3 {
			fmt.Println("参数不足")
			os.Exit(1)
		}
		data:=cmds[2]
		cli.addBlock(data)
	case "printChain":
		fmt.Println("打印区块连命令被调用！")
		cli.printChain()
	default:
		fmt.Println("无效的命令！")
		fmt.Println(Usage)
		os.Exit(1)
	}

}
