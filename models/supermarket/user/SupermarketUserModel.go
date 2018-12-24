package user

import (
	"github.com/flashfeifei/supermarket/models/supermarket"
	"strconv"
	"time"
)

//orm模型
type SupermarketUserModel struct {
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

func (this *SupermarketUserModel) TableName() string {
	return this.Supermarket.TableName() + "user"
}

//生成一个账号
func (this *SupermarketUserModel) CreateUsername() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
