package theme

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/supermarket/theme"
	"github.com/flashfeifei/supermarket/service/supermarket/attachment"
	"time"
)

type themeService struct {
}

func NewThemeService() *themeService {
	return &themeService{}
}

//添加一个主题
func (this *themeService) AddTheme(attachment_id int64, title string) (theme_model *theme.SupermarketThemeModel, err error) {
	image_service := attachment.NewAttachmentImageService()
	_, err = image_service.QueryImageAttachment(attachment_id)
	if err != nil {
		//找不到这个附件
		return nil, err
	}
	o := orm.NewOrm()
	theme_model = new(theme.SupermarketThemeModel)
	theme_model.Title = title
	theme_model.AttachmentId = attachment_id
	theme_model.Createtime = time.Now().Unix()
	theme_model.Updatetime = time.Now().Unix()
	_, err = o.Insert(theme_model)
	if err != nil {
		return nil, err
	}
	return theme_model, nil
}

//更新主题
func (this *themeService) UpdateTheme(id, attachment_id int64, title string) (num int64, err error) {
	image_service := attachment.NewAttachmentImageService()
	_, err = image_service.QueryImageAttachment(attachment_id)
	if err != nil {
		//找不到这个附件
		return 0, err
	}
	o := orm.NewOrm()
	theme_model := new(theme.SupermarketThemeModel)
	theme_model.Id = id
	err = o.Read(theme_model)
	if err != nil {
		return 0, err
	}
	theme_model.AttachmentId = attachment_id
	theme_model.Title = title
	theme_model.Updatetime = time.Now().Unix()
	num, err = o.Update(theme_model)
	if err != nil {
		return 0, err
	}
	return num, nil
}

//软删除一个主题
func (this *themeService) DeleteTheme(id int64) (num int64, err error) {
	o := orm.NewOrm()
	theme_model := new(theme.SupermarketThemeModel)
	theme_model.Id = id
	err = o.Read(theme_model)
	if err != nil {
		return 0, err
	}
	theme_model.Deletetime = time.Now().Unix()
	num, err = o.Update(theme_model)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (this *themeService) Getattachmentlist(page int64, page_size int64, sort string) (themes []orm.Params, count int64) {
	o := orm.NewOrm()
	theme_model := new(theme.SupermarketThemeModel)
	qs := o.QueryTable(theme_model)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Filter("deletetime__exact", 0).Values(&themes, "Id", "Title", "AttachmentId", "Updatetime", "Createtime")
	count, _ = qs.Count()

	//数据格式转化
	image_service := attachment.NewAttachmentImageService()
	attachment_service := attachment.NewAttachmentService()
	for _, b := range themes {
		attachment_id := b["AttachmentId"].(int64)
		att_model, err := image_service.QueryImageAttachment(attachment_id)
		if err != nil {
			continue
		}
		//有外链取外链
		if len(att_model.Links) == 0 {
			b["Url"] = attachment_service.ChangeFilepath(att_model.Filepath)
		} else {
			b["Url"] = att_model.Links
		}
	}
	return themes, count
}
