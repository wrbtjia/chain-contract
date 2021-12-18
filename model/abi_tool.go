package model

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
)

var MasterChef string
var Token string

var MasterChefAbi abi.ABI

func init()  {
	mcData,err:=ioutil.ReadFile("conf/master_chef.txt")
	if err != nil{
		log.Fatal("ReadFile err : ",err)
	}
	MasterChef = string(mcData)

	contractAbi,err:=abi.JSON(bytes.NewBuffer(mcData))
	if err != nil{
		log.Fatal("JSON : ",err)
	}
	MasterChefAbi = contractAbi

	tData,err:=ioutil.ReadFile("conf/token.txt")
	if err != nil{
		log.Fatal("ReadFile err : ",err)
	}
	Token =  string(tData)
}



// 根据函数的名称生成 methodId。abiStr 是智能合约的“abi”数据
func MakeMethoId(methodName string,abiStr string) (string,error) {
	abi := &abi.ABI{} // 实例化 “ABI” 结构体对象指针
	err := abi.UnmarshalJSON([]byte(abiStr))
	if err != nil {
		return "",err
	}
	// 根据 methodName 获取对应的 Method 对象
	method := abi.Methods[methodName]
	methodIdBytes := method.ID // 调用生成 methodId 的函数
	methodId := "0x"+common.Bytes2Hex(methodIdBytes)
	return methodId,nil
}
