package routers

import (
	"SimpleCD/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/webhook", &controllers.MainController{})
}
