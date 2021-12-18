package middleware

import (
	"chain-contract/model"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"strconv"
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


func (e *EthClient)PoolLength(contractAddress string)  (uint64,error) {

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
	return resp[0].(*big.Int).Uint64(),nil
}

func (e *EthClient)PoolInfo(contractAddress string,i int)  ([]interface{},error) {

	methodId,err := model.MakeMethoId("poolInfo",string(model.MasterChef))
	if err != nil{
		log.Println("MakeMethodId err : ",err)
	}
	int64Str := strconv.Itoa(i)
	arg1:= common.HexToHash(int64Str).String()[2:]
	callArg := map[string]interface{}{
		"To": contractAddress,
		"Gas": hexutil.EncodeUint64(300000),
		"Data": methodId+arg1,
	}

	//0000000000000000000000000000000000000000000000000000000000000062
	//0000000000000000000000000000000000000000000000000000000000000059
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


func (e *EthClient)UserInfo(contractAddress string,i uint64,addr string)  ([]interface{},error) {

	methodId,err := model.MakeMethoId("userInfo",string(model.MasterChef))
	if err != nil{
		log.Println("MakeMethodId err : ",err)
	}
	int64Str := strconv.FormatUint(i, 10)
	arg1:= common.HexToHash(int64Str).String()[2:]
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
