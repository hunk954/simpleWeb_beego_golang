package routers

import (
	"testApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/helloWorld", &controllers.MainController{},"get:HelloWorld")
	beego.Router("/login", &controllers.MainController{}, "get:Login")
	beego.Router("/:id([a-zA-Z0-9_]+)", &controllers.MainController{}, "get:LoginIn")
}
