package attachment

import (
	"github.com/flashfeifei/supermarket/lib"
	"github.com/flashfeifei/supermarket/service/supermarket/attachment"
	"strconv"
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
	image_id, image_model, err := image_service.AddImageAttachmentByUpload("imgFile", this.Ctx)
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	att_service := attachment.NewAttachmentService()
	url := att_service.ChangeFilepath(image_model.Filepath)
	result_data := make(map[string]interface{})
	result_data["url"] = url
	result_data["image_id"] = image_id
	this.Ctx.Output.JSON(lib.ApiSuccess(result_data), false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}

//更新图片的信息
func (this *AttachmentController) UpdateImageInfo() {
	field := make(map[string]string)
	field["title"] = this.GetString("Title")
	field["links"] = this.GetString("Links")
	attachment_id, err := strconv.ParseInt(this.GetString("AttachmentId"), 10, 64)
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	image_service := attachment.NewAttachmentImageService()
	_, err = image_service.UpdateImageAttachmentInfo(attachment_id, field)
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	this.Ctx.Output.JSON(lib.ApiSuccess(""), false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}
