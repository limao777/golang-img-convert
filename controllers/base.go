package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
    c.Data["filepath"] = beego.AppConfig.String("filepath")
}

