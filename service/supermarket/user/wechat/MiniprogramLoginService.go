package wechat

import (
	"github.com/astaxie/beego"
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

//创建一个服务
func NewMiniProgramLoginService(Openid, Unionid, SessionKey string) *miniprogramLoginService {
	return &miniprogramLoginService{
		openid:      Openid,
		unionid:     Unionid,
		session_key: SessionKey,
	}
}

//小程序登录
func (this *miniprogramLoginService) Login() (token *miniprogramLoginToken, err error) {

	//获取微信用户信息
	miniprogram_user_model, err := this.register()
	if err != nil {
		return nil, err
	}

	//保存token在缓存
	token = NewMiniprogroamLoginToken(miniprogram_user_model.UserId, miniprogram_user_model.Openid, miniprogram_user_model.Unionid, this.session_key)
	err = token.Save()
	if err != nil {
		return nil, err
	}

	return token, nil
}

//微信小程序用户注册
func (this *miniprogramLoginService) register() (miniprogram_user_model *wechat.MiniprogramUserModel, err error) {
	o := orm.NewOrm()
	//事务开启
	o.Begin()
	//小程序用户模型
	miniprogram_user_model = new(wechat.MiniprogramUserModel)
	//查询构建器
	qs := o.QueryTable(miniprogram_user_model)
	//获取微信小程序的用户
	err = qs.Filter("openid", this.openid).Filter("unionid", this.unionid).Filter("account_type", wechat.USER_TYPE_MINIPROGRAM).
		One(miniprogram_user_model)

	if err == orm.ErrNoRows {
		beego.Debug("找不到找不到找不到好几次")
		//找不到记录
		//注册一下用户
		spermarket_user := new(user.SupermarketUserModel)
		qs := o.QueryTable(spermarket_user)
		account := spermarket_user.CreateUsername()
		err = qs.Filter("username", account).One(spermarket_user)
		if err == orm.ErrNoRows {
			//注册supermarket用户
			spermarket_user.Username = account
			spermarket_user.Password = "123456"
			spermarket_user.Nickname = "萌新托尼"
			spermarket_user.Status = 1
			spermarket_user.Createtime = time.Now().Unix()
			spermarket_user.Updatetime = time.Now().Unix()
			spermarket_user.Deletetime = 0
			user_id, err := o.Insert(spermarket_user)
			if err != nil {
				o.Rollback()
				return nil, err
			}
			//注册微信用户，并绑定supermarket用户
			miniprogram_user_model.AccountType = wechat.USER_TYPE_MINIPROGRAM
			miniprogram_user_model.Openid = this.openid
			miniprogram_user_model.Unionid = this.unionid
			miniprogram_user_model.Createtime = time.Now().Unix()
			miniprogram_user_model.Updatetime = time.Now().Unix()
			miniprogram_user_model.Deletetime = 0
			miniprogram_user_model.UserId = user_id
			_, err = o.Insert(miniprogram_user_model)
			if err != nil {
				//事务回滚
				o.Rollback()
				return nil, err
			}
			//事务提交
			o.Commit()
			return miniprogram_user_model, nil
		} else {
			o.Rollback()
			return nil, err
		}
	}
	o.Commit()
	return miniprogram_user_model, nil
}
