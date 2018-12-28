package banner

import (
	"github.com/flashfeifei/supermarket/lib"
	"github.com/flashfeifei/supermarket/service/supermarket/attachment"
	"github.com/flashfeifei/supermarket/service/supermarket/banner"
	"github.com/pkg/errors"
)

type BannerController struct {
	baseController
}

func (this *BannerController) Index() {
	if this.IsAjax() {
		page, _ := this.GetInt64("page")
		page_size, _ := this.GetInt64("rows")
		sort := this.GetString("sort")
		order := this.GetString("order")
		if len(order) > 0 {
			if order == "desc" {
				sort = "-" + sort
			}
		} else {
			sort = "Id"
		}
		banner_service := banner.NewBannerService()
		users, count := banner_service.Getattachmentlist(page, page_size, sort)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
	} else {
		this.Layout = this.GetTemplate() + "/banner/layout.html"
		this.TplName = this.GetTemplate() + "/banner/banner.html"
	}
}

////添加一张banner图
func (this *BannerController) AddBanner() {
	attachment_id, err := this.GetInt64("AttachmentId")
	title := this.GetString("Title")
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	if len(title) == 0 {
		this.Ctx.Output.JSON(lib.ApiErr(errors.New("title不能为空")), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	_, err = attachment.NewAttachmentImageService().QueryImageAttachment(attachment_id)
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}

	banner_service := banner.NewBannerService()
	_, err = banner_service.AddBanner(attachment_id, title)

	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}

	this.Ctx.Output.JSON(lib.ApiSuccess(""), false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}

//更新banner
//前端技术有限，我这里只更新了title
func (this *BannerController) UpdateBanner() {
	banner_id, err1 := this.GetInt64("Id")
	attachment_id, err2 := this.GetInt64("AttachmentId")
	title := this.GetString("Title")
	if err1 != nil && err2 != nil {
		ers := make(map[string]interface{})
		ers["banner_err"] = err1
		ers["attachment_err"] = err2
		this.Ctx.Output.JSON(lib.ApiErr(ers), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	banner_service := banner.NewBannerService()
	_, err := banner_service.UpdateBanner(banner_id, attachment_id, title)
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	this.Ctx.Output.JSON(lib.ApiSuccess(""), false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}
