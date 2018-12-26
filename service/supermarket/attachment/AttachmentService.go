package attachment

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/supermarket/attachment"
	"io"
	"os"
	"time"
)

type attachmentService struct {
}

func NewAttachmentService() *attachmentService {
	return &attachmentService{}
}

//upload_key是请求协议中，哪个字段保存了图片的信息
//tofile是保存到文件,路径是相对exe的
//剩下的都是数据库字段
func (this *attachmentService) AddImageAttachment(upload_key, tofile string) (error) {
	beego_context := context.NewContext()
	//golang上传文件
	//从网络中读取文件
	file, _, err := beego_context.Request.FormFile(upload_key)
	if err != nil {
		//发生错误
		return err
	}
	//关闭io读写
	defer file.Close()

	//把图片写入文件
	//打开文件
	f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, file)

	if err != nil {
		return err
	}
	o := orm.NewOrm()
	attachment_model := new(attachment.SupermarketAttachmentModel)
	attachment_model.Title = title
	attachment_model.Filepath = tofile
	attachment_model.MimeType = mime_type
	attachment_model.Links = links
	attachment_model.Filetype = attachment.FILE_TYPE_IMAGE
	attachment_model.Createtime = time.Now().Unix()
	attachment_model.Updatetime = time.Now().Unix()
	attachment_model.Deletetime = 0
	o.Insert(attachment_model)
}
