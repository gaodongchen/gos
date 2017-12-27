package routers

import (
	"github.com/astaxie/beego"
	"gos/controllers"
)

func init() {
	beego.Router("/password-generator", &controllers.PasswordGeneratorController{})
	beego.Router("/md5", &controllers.Md5Controller{})
}
