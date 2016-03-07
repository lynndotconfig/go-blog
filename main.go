package main

import (
	_ "goblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	models "goblog/models"
)

func main() {
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		beego.Error(err)
	}
	beego.Run()
}

func init() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
	orm.RegisterModel(new(models.Article))
}
