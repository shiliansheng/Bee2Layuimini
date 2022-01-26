package routers

import (
	"beeappWithLayui/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index.html", &controllers.IndexController{})
	beego.Router("/*", &controllers.GeneralController{})
}
