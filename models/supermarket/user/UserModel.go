package user

import "github.com/flashfeifei/supermarket/models/supermarket"

//orm模型
type User struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int    `orm:"pk;auto"`
	Username                string `orm:"size(32)"`
	Password                string `orm:"size(40)"`
	Nickname                string `orm:"size(10)"`
	Status                  int    `orm:default(1)"`
	Email                   string `orm:"default(null);size(32)"`
	Phone                   string `orm:"default(null);size(20)"`
	Createtime              int64    `orm:"auto_now_add"`
	Updatetime              int64    `orm:"auto_now"`
	Deletetime              int64    `orm:"null"`
}

func (this *User) TableName() string {
	return this.Supermarket.TableName() + "user"
}