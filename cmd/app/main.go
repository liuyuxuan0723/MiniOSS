package main

import (
	"MiniOSS/internal/api"
	"MiniOSS/internal/config"
	"MiniOSS/pkgs"
)

func init() {
	config.LoadConfig()
	pkgs.MysqlInit()
}

func main() {
	r := api.SetRouter()
	r.Run(":8080")
}
