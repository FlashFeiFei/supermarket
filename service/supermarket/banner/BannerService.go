package banner

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/supermarket/banner"
	"github.com/flashfeifei/supermarket/service/supermarket/attachment"
	"time"
)

type bannerService struct {
}

func NewBannerService() *bannerService {
	return &bannerService{}
}

//添加一个banner
func (this *bannerService) AddBanner(attachment_id int64, title string) (banner_model *banner.SupermarketBannerModel, err error) {
	image_service := attachment.NewAttachmentImageService()
	_, err = image_service.QueryImageAttachment(attachment_id)
	if err != nil {
		//找不到这个附件
		return nil, err
	}
	o := orm.NewOrm()
	banner_model = new(banner.SupermarketBannerModel)
	banner_model.Title = title
	banner_model.AttachmentId = attachment_id
	banner_model.Createtime = time.Now().Unix()
	banner_model.Updatetime = time.Now().Unix()
	_, err = o.Insert(banner_model)
	if err != nil {
		return nil, err
	}
	return banner_model, nil
}

//更新banner
func (this *bannerService) UpdateBanner(id, attachment_id int64, title string) (num int64, err error) {
	image_service := attachment.NewAttachmentImageService()
	_, err = image_service.QueryImageAttachment(attachment_id)
	if err != nil {
		//找不到这个附件
		return 0, err
	}
	o := orm.NewOrm()
	banner_model := new(banner.SupermarketBannerModel)
	banner_model.Id = id
	err = o.Read(banner_model)
	if err != nil {
		return 0, err
	}
	banner_model.AttachmentId = attachment_id
	banner_model.Title = title
	banner_model.Updatetime = time.Now().Unix()
	num, err = o.Update(banner_model)
	if err != nil {
		return 0, err
	}
	return num, nil
}

//软删除banner
func (this *bannerService) DeleteBanner(id int64) (num int64, err error) {
	o := orm.NewOrm()
	banner_model := new(banner.SupermarketBannerModel)
	banner_model.Id = id
	err = o.Read(banner_model)
	if err != nil {
		return 0, err
	}
	banner_model.Deletetime = time.Now().Unix()
	num, err = o.Update(banner_model)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (this *bannerService) Getattachmentlist(page int64, page_size int64, sort string) (banners []orm.Params, count int64) {
	o := orm.NewOrm()
	banner_model := new(banner.SupermarketBannerModel)
	qs := o.QueryTable(banner_model)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Filter("deletetime__exact", 0).Values(&banners, "Id", "Title", "AttachmentId", "Updatetime", "Createtime")
	count, _ = qs.Count()

	//数据格式转化
	image_service := attachment.NewAttachmentImageService()
	attachment_service := attachment.NewAttachmentService()
	for _, b := range banners {
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
	return banners, count
}
