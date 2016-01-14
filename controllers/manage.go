package controllers

import (
	"github.com/astaxie/beego"
)

type ManageController struct {
	beego.Controller
}

func (manage *ManageController) Delete() {
	// convert the string value to an int
	articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
}