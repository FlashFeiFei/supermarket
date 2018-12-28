package theme

import (
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/controllers/admin/theme"
)

func AdminThemeRouterRegister() {
	beego.Router("/theme/theme/Index", &theme.ThemeController{}, "*:Index")

	//添加一个主题
	beego.Router("/theme/theme/AddTheme", &theme.ThemeController{}, "*:AddTheme")
	//更新主题
	beego.Router("/theme/theme/UpdateTheme", &theme.ThemeController{}, "*:UpdateTheme")
}
