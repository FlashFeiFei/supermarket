package attachment

import "github.com/flashfeifei/supermarket/models/supermarket"

type SupermarketAttachmentModel struct {
	supermarket.Supermarket `orm:"-"`
	Id                      int64 `orm:"pk;auto"`
	Title                   string
	Filepath                string
	Filetype                int64
	MimeType                string
	Links                   string
	Updatetime              int64
	Createtime              int64
	Deletetime              int64
}

func (this *SupermarketAttachmentModel) TableName() string {
	return this.Supermarket.TableName() + "attachment"
}
