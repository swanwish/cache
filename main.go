package main

import (
	"github.com/liwp-stephen/null/zero"
	_ "github.com/swanwish/go-helper/db"
	_ "github.com/swanwish/go-helper/utils"
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
	//	err := config.Load("config.ini")
	//	if err != nil {
	//		logs.Errorf("Failed to parse configuration, the error is %v", err)
	//		return
	//	}
	//
	//	value, err := config.Get("dbName")
	//	if err != nil {
	//		logs.Errorf("Failed to get db name, the error is %v", err)
	//		return
	//	}
	//	logs.Debugf("The value is %s", value)

	//	logs.Debugf(utils.GetMD5Hash("m0-600-20090405-01.mp3"))
	//	connection := db.DefaultDB{ConnectionGetterFunc: nil}
	//	connection.Exec("")
	type Menu struct {
		MenuId          string      `json:"menuId"`
		MenuName        string      `json:"menuName"`
		MenuDescription zero.String `json:"menuDescription"`
	}
	//	menu := Menu{}
	//	web.Populate(&menu)
	//	logs.Debugf("%s", menu)
}
