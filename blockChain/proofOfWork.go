package main

import (
	"math/big"

	"fmt"
	"crypto/sha256"
	"bytes"
	"math"
)
//定一个工作量证明的结构体
//-block
//-目标值

type ProofOfWork struct {
	//区块数据
	block *Block
	//目标值，先写成固定的值后面在进行推倒演算
	target big.Int
}
//提供一个创建pow的方法
func NewProofOfWork(block *Block)(*ProofOfWork)  {
	//自定义难度值，先写成固定值
	bigIntTmp:=big.NewInt(1)

	bigIntTmp.Lsh(bigIntTmp,256-bits)

	pow:=ProofOfWork{
		block:block,
		target:*bigIntTmp,
	}
	return &pow
}
//提供一个计算哈希值的方法
func (pow *ProofOfWork)Run()( []byte, uint64)  {
	fmt.Printf("挖矿中。。。。。。\n")
	//定义一个nonce变量，用于不断变化
	var nonce uint64
	var hash [32]byte
	for nonce<=math.MaxInt64{
		fmt.Printf("%x\r",hash)
		//1.拿到block数据：pow.block
		hash=sha256.Sum256(pow.prepareData(nonce))
		//需要一个中间变量，将[]byte转换为big.int
		var bigIntTemp big.Int
		bigIntTemp.SetBytes(hash[:])
		//进行判断
		if bigIntTemp.Cmp(&pow.target)==-1 {
			fmt.Printf("挖矿成功,hash：%x,nonce：%d\n",hash,nonce)
			break
		}else {
			nonce++
		}
	}
	return hash[:],nonce
}
func (pow *ProofOfWork)prepareData(nonce uint64)[]byte  {
	b:=pow.block
	tmp:=[][]byte{
		uint2Bytes(b.Version),
		b.PrevBlockHash,
		b.Hash,
		b.MerkelRoot,
		uint2Bytes(b.TimeStamp),
		uint2Bytes(b.Bits),
		uint2Bytes(nonce),
		//b.Data,
		//TODO 美克尔根
	}
	blockInfo:=bytes.Join(tmp,[]byte{})
	return blockInfo
}

