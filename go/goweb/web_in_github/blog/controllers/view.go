package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)
type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.html"
	this.TplName = "view.html"
}
