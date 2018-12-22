package user

import "github.com/flashfeifei/supermarket/models/supermarket"

//orm模型
type SupermarketUser struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
	Username                string
	Password                string
	Nickname                string
	Status                  int
	Email                   string
	Phone                   string
	Createtime              int64
	Updatetime              int64
	Deletetime              int64
}

func (this *SupermarketUser) TableName() string {
	return this.Supermarket.TableName() + "wechat_user"
}
