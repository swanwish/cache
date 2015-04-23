package main

import (
	"github.com/swanwish/go-helper/config"
	"github.com/swanwish/go-helper/logs"
)

func main() {
	//	localIps, err := utils.GetLocalIPAddrs()
	//	if err != nil {
	//		log.Printf("Failed to get local ip address, the error is %v\n", err)
	//		return
	//	}
	//	for _, ip := range localIps {
	//		log.Printf("IP Address: %s\n", ip)
	//	}
	err := config.Load("config.ini")
	if err != nil {
		logs.Errorf("Failed to parse configuration, the error is %v", err)
		return
	}

	value, err := config.Get("dbName")
	if err != nil {
		logs.Errorf("Failed to get db name, the error is %v", err)
		return
	}
	logs.Debugf("The value is %s", value)
}
