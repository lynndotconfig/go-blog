package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
	"goblog/models"
	"fmt"
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

func (manage *ManageController) View() {
	flash := beego.ReadFromRequest(&manage.Controller)
	manage.TplNames = "view.tpl"

	if ok := flash.Data["error"]; ok != "" {
		// Display error message
		manage.Data["error"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	var articles []*models.Article
	num, err := o.QueryTable("Article").All(&articles)

	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = articles
	}
}

func (manage *ManageController) Add() {
	manage.TplNames = "add.tpl"
	o := orm.NewOrm()
	o.Using("default")
	article := models.Article{}

	if err := manage.ParseForm(&article); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		manage.Data["Article"] = article
	}

	if manage.Ctx.Input.Method() == "POST" {
		valid := validation.Validation{}
		isValid, _ := valid.Valid(article)
		if !isValid {
			manage.Data["Error"] = valid.ErrorsMap
			beego.Error("Form didn't validate.", valid.ErrorsMap)
		} else {
			id, err := o.Insert(&article)
			if err == nil {
				msg := fmt.Sprintf("Article inserted with id: ", id)
				beego.Debug(msg)
			} else {
				msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
				beego.Debug(msg)
			}
		}
	}
}