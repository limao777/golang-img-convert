package routers

import (
	"github.com/astaxie/beego"
	"img_convert/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//自带的正则路由还是不太强大，后面看看原生的来处理
	beego.Router("/:pa1([a-zA-Z0-9]+)/:pa2([a-zA-Z0-9]+)/:pa3([0-9]+)/:pa4([0-9]+)/:pa5([0-9]+)/:pa6([a-zA-Z0-9]+)_:pa7([0-9]+)_:pa8([0-9]+).:pa9([a-zA-Z]+)", &controllers.ConvController{})

}
