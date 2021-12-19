package handler

import (
	"chain-contract/middleware"
	"chain-contract/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Process(c *gin.Context)  {


	log.Println("========================")
	contractAddress :="0xDbc1A13490deeF9c3C12b44FE77b503c1B061739"

	client, _ := middleware.Dial("https://bsc-dataseed3.binance.org")


/*	size,err:=client.PoolLength(contractAddress)
	if err != nil{
		log.Println(err)
	}*/
	var i int64 = 0
	for  ;i<95;i++ {

		userResp,err:=client.UserInfo(contractAddress,i,"0xABd99E7d0a07C577Fa397f505Cb4Bd44FAC2c747")
		if err != nil {
			return
		}

		amount:=util.ToDecimal(userResp[0],18)

		log.Println("amount :",amount)

		resp,err := client.PoolInfo(contractAddress,i)
		if err != nil {
			return
		}


		lpAddr :=resp[0].(common.Address)


		symbol,err:=client.Symbol(lpAddr.String())
		decimal,err:=client.Decimal(lpAddr.String())

		log.Println(lpAddr,i,symbol,decimal)
	}

	c.JSON(http.StatusOK, "success")
	return
}
