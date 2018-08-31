package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

type UsersController struct {
	beego.Controller
}

type Users struct{
	Id 		int		`json:"id"`
	Name 	string	`json:"name"`
	Mobile  string	`json:"mobile"`
}

func (c *UsersController) Get() {

	orm.RegisterDataBase(
		"default",
		"mysql",
		beego.AppConfig.String("mysqluser")+
			":"+
			beego.AppConfig.String("mysqlpass")+
			"@tcp(115.159.196.171:3306)/db_name?charset=utf8")


	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	// register model

}



func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//func (c *UserController) Get() {
//	c.Data["website"] = "gaogess"
//	c.Data["Email"] = "gaooooge@gmail.com"
//	c.TplName = "index.tpl"
//}

