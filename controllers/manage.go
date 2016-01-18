package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"goblog/models"
)

type ManageController struct {
	beego.Controller
}

func (manage *ManageController) Delete() {
	// convert the string value to an int
	articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Using("default")
	article := models.Article{}
	// Check if the article exists first
	if exist := o.QueryTable(article.TableName()).Filter("Id", articleId).Exist();exist{
		if num, err := o.Delete(&models.Article{Id: articleId}); err == nil{
			beego.Info("Record Deleted. ", num)
		} else {
			beego.Error("Record couldn't be deleted. Reson: ", err)
		}
	} else {
		beego.Info("Record Doesn't exist.")
	}
}

func (manage *ManageController) Update() {
	o := orm.NewOrm()
	o.Using("default")
	flash := beego.NewFlash()

	//convert the string value to an int
	if articleId, err := strconv.Atoi(manage.Ctx.Input.Param(":id")); err == nil {
		article := models.Article{Id: articleId}
		if o.Read(&article) == nil {
			article.Client = "Sitepoint"
			article.Url = "http:"
			if num, err := o.Update(&article); err == nil {
				flash.Notice("Record Was Update.")
				flash.Store(&manage.Controller)
				beego.Info("Record Was Update. ", num)
			} else {
				flash.Notice("Record Was NOT Updated.")
				flash.Store(&manage.Controller)
				beego.Error("Couldn't find article matching id: ", articleId)
			}
		} else {
				flash.Notice("Record Was NOT Updated.")
				flash.Store(&manage.Controller)
				beego.Error("Couldn't convert id from a string to a number. ", err)

		}
	}

	// redirect afterwards
	manage.Redirect("/manage/view", 302)

}