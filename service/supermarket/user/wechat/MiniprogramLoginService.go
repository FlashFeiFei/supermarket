package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/supermarket/user"
	"github.com/flashfeifei/supermarket/models/supermarket/user/wechat"
	"time"
)

type miniprogramLoginService struct {
	openid      string
	unionid     string
	session_key string
}

//小程序登录
func (this *miniprogramLoginService) Login() (token *miniprogramLoginToken, err error) {
	o := orm.NewOrm()
	//小程序用户模型
	miniprogram_user_model := new(wechat.MiniprogramUserModel)
	//查询构建器
	qs := o.QueryTable(miniprogram_user_model)
	//获取微信小程序的用户
	err = qs.Filter("openid", this.openid).Filter("unionid", this.unionid).Filter("account_type", wechat.USER_TYPE_MINIPROGRAM).
		Filter("deletetime__gt", 0).One(&miniprogram_user_model)

	if err == orm.ErrNoRows {
		//找不到记录
		//注册一下,微信用户模型
		miniprogram_user_model.AccountType = wechat.USER_TYPE_MINIPROGRAM
		miniprogram_user_model.Openid = this.openid
		miniprogram_user_model.Unionid = this.unionid
		now_time := time.Now()
		miniprogram_user_model.Createtime = now_time.Unix()
		miniprogram_user_model.Updatetime = now_time.Unix()
		miniprogram_user_model.Deletetime = 0
		_, err = o.Insert(&miniprogram_user_model)
		if err != nil {
			return nil, err
		}
	}

	//保存token在缓存
	token = NewMiniprogroamLoginToken(miniprogram_user_model.Id, miniprogram_user_model.Openid, miniprogram_user_model.Unionid, this.session_key)
	err = token.Save()
	if err != nil {
		return nil, err
	}

	return token, nil
}

//小程序用户注册
func (this *miniprogramLoginService) Register() (model *wechat.MiniprogramUserModel, err error) {
	o := orm.NewOrm()
	//小程序用户模型
	miniprogram_user_model := new(wechat.MiniprogramUserModel)
	//查询构建器
	qs := o.QueryTable(miniprogram_user_model)
	//获取微信小程序的用户
	err = qs.Filter("openid", this.openid).Filter("unionid", this.unionid).Filter("account_type", wechat.USER_TYPE_MINIPROGRAM).
		Filter("deletetime__gt", 0).One(&miniprogram_user_model)

	if err == orm.ErrNoRows {
		o.Begin()
		//找不到记录
		//注册一下用户
		spermarket_user := new(user.SupermarketUserModel)
		//账号
		spermarket_user.Username = spermarket_user.CreateUsername()
		//密码
		spermarket_user.Password = "123456"
		//昵称
		spermarket_user.Nickname = "萌新托尼"
		spermarket_user.Status = 1
		user_id, err := o.Insert(&spermarket_user)
		if err != nil {
			o.Rollback()
			return nil, err
		}
		miniprogram_user_model.AccountType = wechat.USER_TYPE_MINIPROGRAM
		miniprogram_user_model.Openid = this.openid
		miniprogram_user_model.Unionid = this.unionid
		now_time := time.Now()
		miniprogram_user_model.Createtime = now_time.Unix()
		miniprogram_user_model.Updatetime = now_time.Unix()
		miniprogram_user_model.Deletetime = 0
		miniprogram_user_model.UserId = user_id
		_, err = o.Insert(&miniprogram_user_model)
		if err != nil {
			//事务回滚
			o.Rollback()
			return nil, err
		}

	}
	return miniprogram_user_model, nil
}
