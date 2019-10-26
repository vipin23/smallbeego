package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "VipinHome.com"
	c.Data["Email"] = "vipin@home.com"
	c.TplName = "index.html"
}
