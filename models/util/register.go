package util

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/admin"
	"github.com/flashfeifei/supermarket/models/supermarket/attachment"
	"github.com/flashfeifei/supermarket/models/supermarket/banner"
	"github.com/flashfeifei/supermarket/models/supermarket/theme"
	"github.com/flashfeifei/supermarket/models/supermarket/user"
	"github.com/flashfeifei/supermarket/models/supermarket/user/wechat"
)

//初始化项目的rbac数据库
func RegisterDBAdminModel() {
	orm.RegisterModel(new(admin.Group), new(admin.Node), new(admin.Role), new(admin.User))
}

//初始化supermarker项目数据库
func RegisterDBSupermarket() {
	//supermarket用户模型
	orm.RegisterModel(new(user.SupermarketUserModel))
	//注册微信用户模型
	orm.RegisterModel(new(wechat.MiniprogramUserModel))

	//注册附件模型
	orm.RegisterModel(new(attachment.SupermarketAttachmentModel))
	//注册banner
	orm.RegisterModel(new(banner.SupermarketBannerModel))
	//注册theme
	orm.RegisterModel(new(theme.SupermarketThemeModel))
}

func RegisterDBModel() {
	//注册后台的rbac数据库
	RegisterDBAdminModel()

	//注册supermarker数据库
	RegisterDBSupermarket()
}
