package main

import (
	"awesomeProject/common"
	"awesomeProject/config"
	"awesomeProject/router"
)

func init() {
	config.LoadConfig()
	common.MysqlInit()
}

func main() {
	r := router.SetRouter()
	r.Run(":8080")
}
