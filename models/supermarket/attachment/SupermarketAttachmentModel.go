package attachment

import "github.com/flashfeifei/supermarket/models/supermarket"

type SupermarketAttachmentModel struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
}

func (this *SupermarketAttachmentModel) TableName() string {
	return this.Supermarket.TableName() + "attachment"
}
