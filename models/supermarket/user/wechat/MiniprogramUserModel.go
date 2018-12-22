package wechat

import "github.com/flashfeifei/supermarket/models/supermarket"

//微信用户模型
type MiniprogramUserModel struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
	Openid                  string
	Unionid                 string
	AccountType             int64
	Createtime              int64
	Updatetime              int64
	Deletetime              int64
}

func (this *MiniprogramUserModel) TableName() string {
	return this.Supermarket.TableName() + "wechat_user"
}
