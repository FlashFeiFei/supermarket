package user

import "github.com/flashfeifei/supermarket/models/supermarket"

type MiniprogramUserModel struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
	Openid                  string
	Unionid                 string
	AccountType             int
	Createtime              int
	Updatetime              int
	Deletetime              int
}

func (this *MiniprogramUserModel) TableName() string {
	return this.Supermarket.TableName() + "user"
}
