package attachment

import (
	"github.com/flashfeifei/supermarket/lib"
	"github.com/flashfeifei/supermarket/service/supermarket/attachment"
)

type AttachmentController struct {
	baseController
}

func (this *AttachmentController) Index() {
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
		attachment_service := attachment.NewAttachmentService()
		users, count := attachment_service.Getattachmentlist(page, page_size, sort)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
	} else {
		this.Layout = this.GetTemplate() + "/attachment/layout.html"
		this.TplName = this.GetTemplate() + "/attachment/attachment.html"
	}
}

//上传一张图片
func (this *AttachmentController) UploadImage() {
	image_service := attachment.NewAttachmentImageService()
	image_id, err := image_service.AddImageAttachmentByUpload("imgFile",this.Ctx)
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	this.Ctx.Output.JSON(lib.ApiSuccess(image_id), false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}
