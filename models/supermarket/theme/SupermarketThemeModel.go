package theme

import "github.com/flashfeifei/supermarket/models/supermarket"

type SupermarketThemeModel struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
	AttachmentId            int64
	Title                   string
	Createtime              int64
	Updatetime              int64
	Deletetime              int64
}

func (this *SupermarketThemeModel) TableName() string {
	return this.Supermarket.TableName() + "theme"
}
