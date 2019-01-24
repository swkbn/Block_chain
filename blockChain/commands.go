package main

func (cli*CLI)createBc()  {

	bc:=CreateBlockChain()
	if bc==nil {
		return
	}
	defer bc.db.Close()
}

func (cli *CLI)addBlock(data string)  {

	bc:=GetBlockChain()

	if bc==nil {
		return
	}
	defer bc.db.Close()
	bc.AddBlock(data)

}

func (cli *CLI)printChain()  {
	bc:=GetBlockChain()
	if bc==nil {
		return
	}
	defer bc.db.Close()
	bc.Print2()

}

