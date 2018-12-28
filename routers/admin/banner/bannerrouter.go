package banner

import (
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/controllers/admin/banner"
)

func AdminBannerRouterRegister() {
	beego.Router("/banner/banner/index", &banner.BannerController{}, "*:Index")
}