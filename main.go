package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	_"hello/models"
)

func main() {
	beego.Run()
}

