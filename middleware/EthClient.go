package middleware

import (
	"chain-contract/model"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

type EthClient struct{
	RpcClient *rpc.Client
}

func Dial(url string) (*EthClient,error) {
	rpcClient,err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	client := &EthClient{
		RpcClient: rpcClient,
	}
	return client,nil
}

func (e *EthClient)Call(callArg interface{})  ([]byte,error) {
	var result hexutil.Bytes
	err := e.RpcClient.Call(&result,"eth_call", callArg, "latest")
	if err != nil {
		return nil, errors.New("执行请求call异常:"+err.Error())
	}
	return result ,nil
}


func (e *EthClient)PoolLength(contractAddress string)  (int64,error) {

	methodId,err := model.MakeMethoId("poolLength",string(model.MasterChef))
	if err != nil{
		log.Println("MakeMethodId err : ",err)
		return 0, err
	}
	callArg := map[string]interface{}{
		"To": contractAddress,
		"Data": methodId,
	}
	var result hexutil.Bytes
	errs := e.RpcClient.Call(&result,"eth_call", callArg, "latest")
	if errs !=nil {
		log.Println("call  PoolLength err : ",errs)
		return 0,errs
	}
	resp,err := model.MasterChefAbi.Unpack("poolLength",result)
	if err !=nil {
		log.Println("PoolLength Unpack  err : ",err)
		return 0, err
	}
	return resp[0].(*big.Int).Int64(),nil
}

func (e *EthClient)PoolInfo(contractAddress string,i int64)  ([]interface{},error) {

	methodId,err := model.MakeMethoId("poolInfo",string(model.MasterChef))
	if err != nil{
		log.Println("MakeMethodId err : ",err)
	}

	arg1:=common.BigToHash(big.NewInt(i)).String()[2:]
	callArg := map[string]interface{}{
		"To": common.HexToAddress(contractAddress),
//		"Gas": hexutil.EncodeUint64(300000),
		"Data": methodId+arg1,
	}
	var result hexutil.Bytes
	errs := e.RpcClient.Call(&result,"eth_call", callArg, "latest")
	if errs !=nil {
		log.Println("call  PoolInfo err : ",errs)
		return nil,errs
	}
	resp,err := model.MasterChefAbi.Unpack("poolInfo",result)
	if err !=nil {
		log.Println("PoolLength Unpack  err : ",err)
	}
	return resp,err
}


func (e *EthClient)UserInfo(contractAddress string,i int64,addr string)  ([]interface{},error) {

	methodId,err := model.MakeMethoId("userInfo",string(model.MasterChef))
	if err != nil{
		log.Println("MakeMethodId err : ",err)
	}
	arg1:=common.BigToHash(big.NewInt(i)).String()[2:]
	arg2:= common.HexToHash(addr).String()[2:]
	callArg := map[string]interface{}{
		"To": contractAddress,
		"Data": methodId+arg1+arg2,
	}
	var result hexutil.Bytes
	errs := e.RpcClient.Call(&result,"eth_call", callArg, "latest")
	if errs !=nil {
		log.Println("call  PoolInfo err : ",errs)
		return nil,errs
	}
	resp,err := model.MasterChefAbi.Unpack("userInfo",result)
	if err !=nil {
		log.Println("PoolLength Unpack  err : ",err)
	}
	return resp,err
}


func (e *EthClient)Symbol(contractAddress string)  (string,error) {

	methodId,err := model.MakeMethoId("symbol",string(model.LpToken))
	if err != nil{
		log.Println("MakeMethodId err : ",err)
	}

	callArg := map[string]interface{}{
		"To": contractAddress,
		"Data": methodId,
	}
	var result hexutil.Bytes
	errs := e.RpcClient.Call(&result,"eth_call", callArg, "latest")
	if errs !=nil {
		log.Println("call  Symbol err : ",errs)
		return "",errs
	}
	resp,err := model.LpTokenAbi.Unpack("symbol",result)
	if err !=nil {
		log.Println("Symbol Unpack  err : ",err)
	}
	return resp[0].(string),err
}

func (e *EthClient)Decimal(contractAddress string)  (uint8,error) {

	methodId,err := model.MakeMethoId("decimals",string(model.LpToken))
	if err != nil{
		log.Println("MakeMethodId err : ",err)
	}

	callArg := map[string]interface{}{
		"To": contractAddress,
		"Data": methodId,
	}
	var result hexutil.Bytes
	errs := e.RpcClient.Call(&result,"eth_call", callArg, "latest")
	if errs !=nil {
		log.Println("call  Decimal err : ",errs)
		return 0,errs
	}
	resp,err := model.LpTokenAbi.Unpack("decimals",result)
	if err !=nil {
		log.Println("Symbol Decimal  err : ",err)
	}
	return resp[0].(uint8),err
}




type CallArg struct {
	// common.Address 是以太坊依赖包的地址类型，其原型是 [20]byte 数组
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gas_price"`
	Value    string `json:"value"`
	Data     string `json:"data"`		// 这个就是 data
	Nonce    string `json:"nonce"`
}
