package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/supermarket/user/wechat"
	"time"
)

type miniprogramLoginService struct {
	openid      string
	unionid     string
	session_key string
}

func (this *miniprogramLoginService) Login() error {
	o := orm.NewOrm()
	//小程序用户模型
	miniprogram_user_model := new(wechat.MiniprogramUserModel)
	//查询构建器
	qs := o.QueryTable(miniprogram_user_model)
	//获取微信小程序的用户
	err := qs.Filter("openid", this.openid).Filter("unionid", this.unionid).Filter("account_type", wechat.USER_TYPE_MINIPROGRAM).
		Filter("deletetime__gt", 0).One(&miniprogram_user_model)

	if err == orm.ErrNoRows {
		//找不到记录
		//注册一下
		miniprogram_user_model.AccountType = wechat.USER_TYPE_MINIPROGRAM
		miniprogram_user_model.Openid = this.openid
		miniprogram_user_model.Unionid = this.unionid
		now_time := time.Now()
		miniprogram_user_model.Createtime = now_time.Unix()
		miniprogram_user_model.Updatetime = now_time.Unix()
		miniprogram_user_model.Deletetime = 0
		id, err := o.Insert(&miniprogram_user_model)
		if err != nil {
			return err
		}
	}

	return nil
}
