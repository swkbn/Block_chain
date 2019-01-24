package main

import (
	"time"

	"bytes"
	"encoding/binary"

	"encoding/gob"
)

const bits  = 16

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
		Version:01,
		PrevBlockHash:prevBlockHash,
		Hash:nil,
		MerkelRoot:nil,
		TimeStamp:uint64(time.Now().Unix()),
		Bits:bits,
		Nonce:0,
		Data:[]byte(data),
	}
	//设置哈希值
	//block.setHash()
	pow:=NewProofOfWork(&block)
	hash,nonce:=pow.Run()
	block.Hash=hash
	block.Nonce=nonce
	return &block

}
//生成哈希
/*func (b*Block)setHash()  {

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

}*/
//将数字转成字节流

func uint2Bytes(num uint64)[]byte  {

	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)

	if err!=nil {

		panic(err)
	}
	return buffer.Bytes()

}


//序列化，将结构体转换为字节流

func (b*Block)Serialize()[]byte  {
	//定义容器
	var buffer bytes.Buffer
	//编码
	//1创建编码器
	encoder:=gob.NewEncoder(&buffer)
	//2编码器encode方法，得到字节流
	err:=encoder.Encode(b)
	//错误校验
	if err!=nil {

		panic(err)
	}

	return buffer.Bytes()
}

//反序列化，将字节流转成block结构体

func Deserialize(data []byte)*Block  {

	var block Block
	//1解码器
	decoder:=gob.NewDecoder(bytes.NewReader(data))
	//解码器decode方法，得到结构体
	err:=decoder.Decode(&block)
	if err!=nil {
		panic(err)
	}
	return &block
}

