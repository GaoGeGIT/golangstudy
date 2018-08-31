package routers

import (
	"green-go/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Get("hello",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

    beego.Router("users", &controllers.UsersController{})

}
