package user

import (
	"github.com/flashfeifei/supermarket/models/supermarket"
	"time"
)

//orm模型
type User struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int       `orm:"pk;auto"`
	Username                string    `orm:"size(32)"`
	Password                string    `orm:"size(40)"`
	Nickname                string    `orm:"size(10)"`
	Status                  int       `orm:default(1)"`
	Email                   string    `orm:"default(null);size(32)"`
	Phone                   string    `orm:"default(null);size(20)"`
	Createtime              time.Time `orm:"auto_now_add;type(int)"`
	Updatetime              time.Time `orm:"auto_now;type(int)"`
	Deletetime              time.Time `orm:"default(0);type(int)"`
}

func (this *User) TableName() string {
	return this.Supermarket.TableName() + "user"
}
