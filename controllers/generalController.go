package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type GeneralController struct {
	beego.Controller
}

func (c *GeneralController) Get() {
	c.TplName = c.Ctx.Request.URL.Path[1:]
}