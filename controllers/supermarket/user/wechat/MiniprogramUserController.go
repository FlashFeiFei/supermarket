package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"net/http"
)

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
	//获取小程序的code
	wx_session, err := this.getWxCode2Session(appid, secret, code)

	if err != nil {
		beego.Error(err.Error())
		this.Data["json"] = err.Error()
		this.Ctx.Output.JSON(wx_session, false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	beego.Debug(wx_session["errcode"])
	//请求接口成功，异常结果
	if wx_session["errcode"] != 0 {
		this.Ctx.Output.JSON(wx_session, false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	beego.Debug("??????????????????????????????")
	//正常结果
	log.Println(wx_session["session_key"])
	log.Println(wx_session["unionid"])
	log.Println(wx_session["openid"])
	this.Ctx.Output.JSON(wx_session, false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}

//发送请求
func (this *MiniprogramUserController) getWxCode2Session(appid, secret, code string) (wx_session map[string]interface{}, err error) {
	client := http.Client{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, code)
	response, err := client.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	//读取网络源数据
	//reselt_string, _ := ioutil.ReadAll(response.Body)
	//log.Println(string(reselt_string))
	//解码
	json_decode := json.NewDecoder(response.Body)
	err = json_decode.Decode(&wx_session)
	if err != nil {
		return nil, err
	}
	beego.Debug(wx_session)
	return wx_session, nil
}
