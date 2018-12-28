package banner

import "github.com/flashfeifei/supermarket/models/supermarket"

type SupermarketBannerModel struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
	Attachmentid            int64
	Title                   string
	Createtime              int64
	Updatetime              int64
	Deletetime              int64
}

func (this *SupermarketBannerModel) TableName() string {
	return this.Supermarket.TableName() + "banner"
}
