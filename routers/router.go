package routers

import (
	"beans/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.NgAppController{})
    beego.Router("/auth", &controllers.AuthController{}, "get:Index")
    beego.Router("/auth/login", &controllers.AuthController{}, "post:Login")
    beego.Router("/auth/logout", &controllers.AuthController{}, "post:Logout")
    beego.Router("/auth/refresh", &controllers.AuthController{}, "get:Refresh")
}
