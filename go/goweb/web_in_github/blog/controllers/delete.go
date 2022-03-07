package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)
type DeleteController struct {
	beego.Cotroller
}

func (this *DeleteController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	blog := models.GetBlog(id)
	this.Data["Post"] = blog
	models.DelBlog(blog)
	this.Ctx.Redirect(302, "/")
}
