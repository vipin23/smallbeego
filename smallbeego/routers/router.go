package routers

import (
	"github.com/vipin23/smallbeego/smallbeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
