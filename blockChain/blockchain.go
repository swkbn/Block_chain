package main

import (
	"fmt"
	"Block_chain/blockChain/lib/bolt"
	"os"
	"time"
)
const genesisInfo = "hello！"

//数据库名
const blockChainFile  = "blockChain.db"
//通名
const blockBucket  = "blockBucket"
//最后一个区块的标识
const lastHashKey  = "lastHashKey"
//引入区块
//使用数组来模拟区块连
type BlockChain struct {
	//数据库句柄
	db*bolt.DB
	//存储最后一个区块哈希
	tail []byte
}
//创建区块连
/*主要作用就是创建区块连并且，添加创世区块*/
func CreateBlockChain()*BlockChain  {
	//判断是否由此文件
	if isFileExist(blockChainFile) {
		fmt.Println("区块已经存在,无需重复创建！")
		return nil
	}
	//创建区块连结构,一般会在创建的时候，添加一个区块，被称之为创世区块
	var bc BlockChain
	//操作数据库
	//创建db文件
	db,err:=bolt.Open(blockChainFile,0600,nil)

	if err!=nil {
		panic(err)
	}
	//2、操作数据库
	db.Update(func(tx *bolt.Tx) error {
		//创建一个桶
		b:=tx.Bucket([]byte(blockBucket))
		if b==nil {
			//A:第一次调用这个方法，里面没有bucket时，需要创建，并且添加创世区块
			b,err:=tx.CreateBucket([]byte(blockBucket))
			if err!=nil {
				panic(err)
			}
			//将创世块写入bucket中
			genesisBlock:=NewBlock(genesisInfo,[]byte{})
			//更新区块2更新lastHashKey的值
			b.Put(genesisBlock.Hash,genesisBlock.Serialize()/*区块的字节流*/)
			b.Put([]byte(lastHashKey),genesisBlock.Hash)
			bc.tail=genesisBlock.Hash
		}
		return nil
	})
	bc.db = db
	return &bc
}

func GetBlockChain()*BlockChain  {
	if !isFileExist(blockChainFile) {
		fmt.Println("请先创建区块连文件")
		return nil
	}
	var bc BlockChain
	//打开数据库
	//创建blockChain.db文件
	db,err:=bolt.Open(blockChainFile,0600,nil)
	if err!=nil {
		panic(err)
	}
	//2.操作数据库
	db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		if b==nil {
			fmt.Println("获取区块连实力时，bucket不应为空！")
			os.Exit(1)
		}
		lastHash:=b.Get([]byte(lastHashKey))
		bc.tail=lastHash
		return nil
	})
	bc.db=db
	return &bc
}
//添加区块
func (bc *BlockChain)AddBlock(data string)bool  {
	//创建一个区块，前区块的哈希值从bc的最后一个块元素即可获取
	/*lasBlock :=bc.Blocks[len(bc.Blocks)-1]//获取最后一个区块
	//即将添加的区块的前哈希值，就是bc中的最后区块的Hahs字段的值
	prevHash:=lasBlock.Hash
	//创建一个新的区块
	newBlock:=NewBlock(data,prevHash)
	//append到区块莲的Blocks数组中
	bc.Blocks=append(bc.Blocks,newBlock)*/
	err:=bc.db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		if b==nil {

			fmt.Println("addBlock时不应该为空")
			os.Exit(1)
		}
		prevHash:=bc.tail
		newBlock:=NewBlock(data,prevHash)
		//更新区块ls
		// ，更新lastHashKey的值
		err:=b.Put(newBlock.Hash,newBlock.Serialize()/*区块的字节流*/)
		if err!=nil {
			return err
		}
		//添加  补充
		err=b.Put([]byte(lastHashKey),newBlock.Hash)
		if err!=nil {
			return err
		}

		//把新的区块放入尾巴中
		bc.tail=newBlock.Hash
		return nil
	})
	if err!=nil {
		return false
	}
	return true
}
//使用bolt字节的ForEach来打印区块连

func (bc*BlockChain)print1()  {

	bc.db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%x\n",k)
			return nil
		})
		return nil
	})
}
//使用我们自己实现的迭代器便利区块链

//使用我们自己实现的迭代器遍历区块链
func (bc *BlockChain) Print2() {
	//1. 创建一个迭代器
	it := bc.NewIterator()
	//3. 终止条件，当前区块的前哈希为空(nil)，跳出循环
	for {
		//2. 使用for循环不断的调用Next方法，得到block，打印
		// 获取区块，指针左移
		block := it.Next()

		//fmt.Printf(" ========= 区块高度 : %d =======\n", i)
		fmt.Printf(" ===============================\n")

		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevBlockHash)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)

		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)

		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳: %s\n", timeFormat)
		//fmt.Printf("时间戳: %d\n", block.TimeStamp)

		fmt.Printf("难度值: %d\n", block.Bits)
		fmt.Printf("随机数: %d\n", block.Nonce)

		pow := NewProofOfWork(block)
		//fmt.Printf("IsValid : %v\n", pow.IsValid())
		fmt.Println(pow)
		fmt.Printf("区块数据: %s\n", block.Data)

		//终止条件，当前区块的前哈希为空(nil)
		//if block.PrevBlockHash == nil
		//if bytes.Equal(block.PrevBlockHash, []byte{}) {

		if len(block.PrevBlockHash) == 0 {
			fmt.Printf("区块链遍历完成!\n")
			break
		}
	}
}
//定义一个专门的迭代器结构，用于遍历区块链
type Iterator struct {
	//1.  db
	db *bolt.DB

	//2.  currentHash
	// - 最开始指向lastHash, 内次调用Next之后，都会向左移动
	currentHash []byte
}

//创建一个迭代器
func (bc *BlockChain) NewIterator() *Iterator {
	it := Iterator{
		db:          bc.db,
		currentHash: bc.tail,
	}

	return &it
}

//实现遍历的方法，Next
func (it *Iterator) Next() *Block {
	var block Block
	//1. db.View
	it.db.View(func(tx *bolt.Tx) error {

		//2. 获取bucket
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			fmt.Printf("遍历区块链时，bucket不应为空!")
			os.Exit(1)
		}

		//3. 获取区块数据，Get（currentHash）
		blockBytesInfo /*block的字节流*/ := b.Get(it.currentHash)

		//反序列化
		block = *Deserialize(blockBytesInfo)

		return nil
	})

	//4. currentHash向左移动, 一定要记得更新
	it.currentHash = block.PrevBlockHash

	return &block
}




