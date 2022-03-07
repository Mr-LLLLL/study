package controllers

import "github.com/astaxie/beego"
type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.html"
	this.TplName = "index.html"
}
