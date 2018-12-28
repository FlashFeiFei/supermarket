package theme

import "github.com/flashfeifei/supermarket/models/supermarket"

type SupermarketThemeModel struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
	Attachmentid            int64
	Title                   string
	createtime              int64
	updatetime              int64
	deletetime              int64
}

func (this *SupermarketThemeModel) TableName() string {
	return this.Supermarket.TableName() + "theme"
}
