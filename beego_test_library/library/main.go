package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "library/models"
	_ "library/routers"
)

func main() {
	orm.Debug = true
	beego.Run()
}

