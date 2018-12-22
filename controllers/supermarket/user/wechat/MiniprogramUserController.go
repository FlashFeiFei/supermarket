package wechat

import "github.com/astaxie/beego"
import "github.com/astaxie/beego/httplib"

type MiniprogramUserController struct {
	beego.Controller
}

//微信小程序获取code
func (this *MiniprogramUserController) Login() {
	//获取小程序传过来的code
	code := this.GetString("code")
	//获取小程序配置文件
	appid := beego.AppConfig.String("wechat_miniprogram_appid")
	secret := beego.AppConfig.String("wechat_miniprogram_secret")

}
