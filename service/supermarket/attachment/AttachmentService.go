package attachment

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	am "github.com/flashfeifei/supermarket/models/supermarket/attachment"
)

type attachmentService struct {
	domain string
}

func NewAttachmentService() *attachmentService {
	return &attachmentService{
		domain: beego.AppConfig.String("supermarket_domain"),
	}
}
func (this *attachmentService) Getattachmentlist(page int64, page_size int64, sort string) (attachments []orm.Params, count int64) {
	o := orm.NewOrm()
	attachment_model := new(am.SupermarketAttachmentModel)
	qs := o.QueryTable(attachment_model)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&attachments, "Id", "Title", "Filepath", "Filetype",
		"Links", "Updatetime", "Createtime")
	count, _ = qs.Count()
	for _, m := range attachments {
		m["Filepath"] = this.ChangeFilepath(m["Filepath"].(string))
		m["Filetype"] = this.ChangeFileType(m["Filetype"].(int64))
	}
	return attachments, count
}

//转化能成能访问的url
func (this *attachmentService) ChangeFilepath(filepath string) (string) {
	return this.domain + filepath
}

func (this *attachmentService) ChangeFileType(file_type int64) (fileType string) {
	switch file_type {
	case am.FILE_TYPE_IMAGE:
		fileType = "图片"
	default:
		fileType = "未知"
	}
	return fileType
}
