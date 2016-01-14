package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (main *MainController) GoBlog() {
	main.Data["Website"] = "My Go Blog"
	main.Data["Email"] = "lynn.config@gmail.com"
	main.Data["EmailName"] = "Liping Wang"
	main.TplNames = "go-blog.tpl"
}