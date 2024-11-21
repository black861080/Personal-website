package main

import (
	"black/router"
	"black/util"
)

func main() {
	util.InitConfig()
	util.InitMysql()
	r := router.Router()
	r.Run()
}
