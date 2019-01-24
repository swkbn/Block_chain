package main

import (
	"bytes"
	"encoding/gob"
	"crypto/sha256"
	//"text/tabwriter"
)

//交易输入

//指明交易发起人可支付金的来源，包含

type TXInput struct {
	//引用utxo所在交易的ID（知道在那个房间）
	TXID []byte

	//所消费utxo在output中的索引（具体位置）
	Index int64
	//解锁脚本（签名，公钥）
	ScriptSig string//先用一个字符串来代替，后续使用函数代替

}

//交易输出

//包含资金接收方的相关信息

type TXOutput struct {
	//接收金额（数字）
	Value float64

	//锁定脚本（对方公钥的哈希，这个哈希可以通过地址反推出来。所以转账时知道地址即可）

	ScriptPubKey string
}

type Transaction struct {

	TxId []byte		//交易id
	TXInputs []TXInput	//输入数组
	TXOutputs []TXOutput	//输入数组
	TimeStamp uint64		//时间戳
}
//交易ID

//一般是交易结构的哈希值（参考block的哈希做法）

func (tx* Transaction)SetTXHash()  {

	//交易id我们使用sha256来获取
	//获取tx字节流

	var buffer bytes.Buffer

	//编码
	//创建编码器
	encoder:=gob.NewEncoder(&buffer)
	//编码器得到字节流

	err:=encoder.Encode(tx)

	if err!=nil {
		panic(err)
	}
	//使用sha256方法
	hash:=sha256.Sum256(buffer.Bytes())

	tx.TxId=hash[:]

}
//挖矿或得金额
const reward  = 12.5

//挖矿交易
func NewCinbaseTX(data string,miner /*矿工地址，使用string代替*/string)*Transaction  {
	txinput:=TXInput{
		TXID:nil,
		Index:-1,
		ScriptSig:data,
	}
	txoutput:=TXOutput{
		Value:reward,
		ScriptPubKey:miner,
	}
	tx:=Transaction{
		TxId:nil,
		TXInputs:[]TXInput{txinput},
		TXOutputs:[]TXOutput{txoutput},
	}
	//调用setTXHasha
	tx.SetTXHash()
	return &tx
}


