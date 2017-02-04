package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
    if c.Data["error_msg"] == nil{
        c.Data["error_msg"] = "404 not found"
    }
    c.TplName = "errors/404.tpl"
}

