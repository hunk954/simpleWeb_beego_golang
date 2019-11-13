package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func(main *MainController) Login(){
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your name"
	main.TplName = "login.tpl"
}

func(main *MainController) HelloWorld(){
	main.Data["Website"] = "HelloWorld Website"
	main.Data["Email"] = "415391519@qq.com"
	main.Data["EmailName"] = "Wang Junhuan"
	main.TplName = "helloWorld.tpl"
}

func(main *MainController) LoginIn(){
	
	main.TplName = "helloWorld.tpl"
}



