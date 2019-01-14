package main

import (
	"time"
	"crypto/sha256"
	"bytes"
	"encoding/binary"
)

//定义结构体（区块头的字段比平常的少）
type Block struct {
	//版本号
	Version uint64


	//前区块哈希
	PrevBlockHash []byte
	//当前区块哈希，这是为了方便加入的字段，正常区块中没有这个字段
	Hash []byte
	//梅克尔根
	MerkelRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Bits uint64
	//随机数(挖矿要求的值)
	Nonce uint64
	//数据
	Data []byte
}

//创建区块

func NewBlock(data string,prevBlockHash []byte) *Block {
	block:=Block{
		Version:00,
		PrevBlockHash:prevBlockHash,
		Hash:nil,
		MerkelRoot:nil,
		TimeStamp:uint64(time.Now().Unix()),
		Bits:13,
		Nonce:0,
		Data:[]byte(data),
	}
	//设置哈希值
	block.setHash()
	return &block

}
//生成哈希
func (b*Block)setHash()  {

	var blockInfo  []byte
	blockInfo=append(blockInfo,uint2Bytes(b.Version)...)
	blockInfo=append(blockInfo,b.PrevBlockHash...)
	blockInfo=append(blockInfo,b.Hash...)


	blockInfo=append(blockInfo,b.MerkelRoot...)
	blockInfo=append(blockInfo,uint2Bytes(b.TimeStamp)...)
	blockInfo=append(blockInfo,uint2Bytes(b.Bits)...)
	blockInfo=append(blockInfo,uint2Bytes(b.Nonce)...)


	blockInfo=append(blockInfo,b.Data...)
	hash:=sha256.Sum256(blockInfo)

	b.Hash=hash[:]

}
//将数字转成字节流

func uint2Bytes(num uint64)[]byte  {

	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)

	if err!=nil {

		panic(err)
	}
	return buffer.Bytes()

}


