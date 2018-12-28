package banner

import (
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/controllers/admin/banner"
)

func AdminBannerRouterRegister() {
	beego.Router("/banner/banner/Index", &banner.BannerController{}, "*:Index")
	//添加banner
	beego.Router("/banner/banner/AddBanner", &banner.BannerController{}, "*:AddBanner")
	//更新banner
	beego.Router("/banner/banner/UpdateBanner", &banner.BannerController{}, "*:UpdateBanner")
}