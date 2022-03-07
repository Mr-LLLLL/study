package controllers

import (
	"time"

	"github.com/astaxie/beego"
)
type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.Layout = "layout.html"
	this.TplName = "new.html"
}

func (this *NewController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now()
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
