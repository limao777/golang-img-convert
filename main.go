package main

import (
	_ "img_convert/routers"
	"img_convert/controllers"
	"github.com/astaxie/beego"
)

func main() {
    beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

