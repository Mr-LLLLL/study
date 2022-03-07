package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var intcount int
	sess := this.StartSession()
	count := sess.Get("count")
	if count == nil {
		intcount = 0
	} else {
		intcount = count.(int)
	}
	intcount += 1
	sess.Set("count", intcount)
	this.Data["Username"] = "astaxie"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.html"
}
