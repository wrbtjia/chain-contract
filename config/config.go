package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init(fileName string,key string)  {

	viper.SetConfigName(fileName)
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	if err:=viper.ReadInConfig();err !=nil{
		log.Println("加载配置文件失败:",err)
	}

	log.Println( viper.GetString("key.userInfo.method"))
	log.Println( viper.GetString("key.userInfo.val"))
	log.Println( viper.GetString("key.poolInfo"))



}

