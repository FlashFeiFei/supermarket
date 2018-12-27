package attachment

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/lib"
	am "github.com/flashfeifei/supermarket/models/supermarket/attachment"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	UPLOAD_IMAGE_PATH = "static/upload/"
)

type attachmentImageService struct {
}

func NewAttachmentImageService() *attachmentImageService {
	return &attachmentImageService{}
}

//通过上传添加一张图片
//upload_key是请求协议中，哪个字段保存了图片的信息
//剩下的都是数据库字段
func (this *attachmentImageService) AddImageAttachmentByUpload(upload_key string) (int64, error) {
	beego_context := context.NewContext()
	//golang上传文件
	//从网络中读取文件
	file, _, err := beego_context.Request.FormFile(upload_key)
	if err != nil {
		//发生错误
		return 0, err
	}
	//关闭io读写
	defer file.Close()

	//把图片写入文件
	//打开文件
	tofile := UPLOAD_IMAGE_PATH + strconv.FormatInt(lib.GetUid(), 10)
	f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	_, err = io.Copy(f, file)

	if err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	attachment_model := new(am.SupermarketAttachmentModel)
	attachment_model.Filepath = tofile
	attachment_model.Filetype = am.FILE_TYPE_IMAGE
	attachment_model.Createtime = time.Now().Unix()
	attachment_model.Updatetime = time.Now().Unix()
	attachment_model.Deletetime = 0
	attchment_id, err := o.Insert(attachment_model)
	if err != nil {
		return 0, err
	}
	return attchment_id, nil
}

//通过外链添加一张图片
func (this *attachmentImageService) AddImageAttachmentByLinks(links string) (int64, error) {
	o := orm.NewOrm()
	attachment_model := new(am.SupermarketAttachmentModel)
	attachment_model.Links = links
	attachment_model.Filetype = am.FILE_TYPE_IMAGE
	attachment_model.Createtime = time.Now().Unix()
	attachment_model.Updatetime = time.Now().Unix()
	attachment_model.Deletetime = 0
	attchment_id, err := o.Insert(attachment_model)
	if err != nil {
		return 0, err
	}
	return attchment_id, nil
}

//更新一下属性
func (this *attachmentImageService) UpdateImageAttachmentInfo(id int64, field map[string]string) (num int64, err error) {
	o := orm.NewOrm()
	attachment_model := new(am.SupermarketAttachmentModel)
	attachment_model.Id = id
	err = o.Read(attachment_model)
	if err != nil {
		return 0, err
	}
	if value, ok := field["title"]; ok {
		attachment_model.Title = value
	}
	if value, ok := field["links"]; ok {
		attachment_model.Title = value
	}
	attachment_model.Updatetime = time.Now().Unix()
	num, err = o.Update(attachment_model)
	if err != nil {
		return 0, err
	}
	return num, nil
}

//删除一个文件
func (this *attachmentImageService) DeleteImageAttachment(id int64) (num int64, err error) {
	o := orm.NewOrm()
	attachment_model := new(am.SupermarketAttachmentModel)
	attachment_model.Id = id
	err = o.Read(attachment_model)
	if err != nil {
		return 0, err
	}
	//删除文件
	tofile := attachment_model.Filepath
	os.Remove(tofile)
	num, err = o.Delete(attachment_model)
	if err != nil {
		return 0, err
	}
	return num, err
}
