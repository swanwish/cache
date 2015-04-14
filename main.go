package main

import (
	"log"

	"github.com/swanwish/go-helper/utils"
)

func main() {
	localIps, err := utils.GetLocalIPAddrs()
	if err != nil {
		log.Printf("Failed to get local ip address, the error is %v\n", err)
		return
	}
	for _, ip := range localIps {
		log.Printf("IP Address: %s\n", ip)
	}
}
