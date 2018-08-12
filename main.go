package main

import (
	"github.com/josimar-jr/go_mongo_expenses_sample/models"
	_ "github.com/josimar-jr/go_mongo_expenses_sample/routers"

	"github.com/astaxie/beego"
)

func initDB() {
	// initiates all the collections in mongo
	models.StartUserCollection()
	models.StartAccountCollection()
	models.StartCategoryCollection()
	models.StartMovementCollection()
}

func main() {
	initDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
