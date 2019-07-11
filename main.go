package main

import (
	"beans/models"
	_ "beans/routers"
	"fmt"
	"github.com/astaxie/beego"
	"os"
)

const VERSION = "0.1.1"

func main() {
	//initialize()
	beego.Run()
}

func init() {
	fmt.Println(beego.AppConfig.String("database::db_name"))
	initialize()
}
//func Run()  {
//	initialize()
//}

func initialize() {

	initArgs()

	models.Connect()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}