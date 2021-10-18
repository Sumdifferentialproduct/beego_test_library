package routers

import (
	"github.com/astaxie/beego"
	"library/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego_test_library
    beego.Router("/book/list", &controllers.BookController{},"*:List")
    beego.Router("/book/delete", &controllers.BookController{},"*:Delete")
    beego.Router("/book/add", &controllers.BookController{},"*:Add")
    beego.Router("/book/edit", &controllers.BookController{},"*:Edit")
}
