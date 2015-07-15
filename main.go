package main

import (
	_ "github.com/swanwish/go-helper/db"
	"github.com/swanwish/go-helper/logs"
	"github.com/swanwish/go-helper/utils"
	_ "github.com/swanwish/go-helper/web"
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
	//	type Menu struct {
	//		MenuId          string      `json:"menuId"`
	//		MenuName        string      `json:"menuName"`
	//		MenuDescription zero.String `json:"menuDescription"`
	//	}
	//	menu := Menu{}
	//	logs.Debug("type: ", reflect.TypeOf(menu))
	//	logs.Debug("value: ", reflect.ValueOf(menu))
	//	v := reflect.ValueOf(&menu)
	//	logs.Debug("type:", v.Type())
	//	logs.Debug("Kind:", v.Kind())
	//	logs.Debug("Setable:", v.CanSet())
	//
	//	ve := v.Elem()
	//	logs.Debug("ve setable:", ve.CanSet())
	//	vet := ve.Type()
	//	for i := 0; i < ve.NumField(); i++ {
	//		f := ve.Field(i)
	//		switch f.Type() {
	//		case "int":
	//			f.SetInt(60)
	//		case "string":
	//			f.SetString("test")
	//		}
	//		logs.Debugf("%d: %s %s = %v\n", i, vet.Field(i).Name, f.Type(), f.Interface())
	//	}

	//	var x int64 = 64
	//	xv := reflect.ValueOf(&x)
	//	logs.Debug("xv type:", xv.Type())
	//	logs.Debug("xv setable:", xv.CanSet())
	//
	//	xve := xv.Elem()
	//	logs.Debug("xve setable:", xve.CanSet())
	//
	//	xve.SetInt(100)
	//	logs.Debug("x:", x)

	//		web.Populate(&menu)
	//	logs.Debugf("%s", menu)
	pwd := utils.GeneratePassword("liwp", 0, 8)
	logs.Debugf("The password is %v", pwd)
	//	seed := "liwp@Stephen我"
	//	var offset int64
	//	for i := 0; i < len(seed); i++ {
	//		offset += int64(seed[i])
	//	}
	//	logs.Debugf("The offset is %d", offset)
}
