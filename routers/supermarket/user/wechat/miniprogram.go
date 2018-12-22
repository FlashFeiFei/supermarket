package wechat

import (
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/controllers/supermarket/user/wechat"
)

func MiniprogramLoginRouter() {
	//小程序获取
	beego.Router("/supermarket/user/wechat/miniprogram/login", &wechat.MiniprogramUserController{}, "*:Login")
}
