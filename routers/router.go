package routers

import (
	"beans/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.NgAppController{})
}
