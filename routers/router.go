package routers

import (
	"goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/blog", &controllers.MainController{}, "get:GoBlog")
    beego.Router("/blog/:id([0-9])", &controllers.MainController{}, "get:GoBlog")
    beego.Router("/manage/delete/:id([0-9])", &controllers.MainController{}, "*:Delete")
}
