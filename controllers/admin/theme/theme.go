package theme

import (
	"github.com/flashfeifei/supermarket/lib"
	"github.com/flashfeifei/supermarket/service/supermarket/attachment"
	"github.com/flashfeifei/supermarket/service/supermarket/theme"
	"github.com/pkg/errors"
)

type ThemeController struct {
	baseController
}

func (this *ThemeController) Index() {
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
		theme_service := theme.NewThemeService()
		themes, count := theme_service.Getattachmentlist(page, page_size, sort)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &themes}
		this.ServeJSON()
	} else {
		this.Layout = this.GetTemplate() + "/theme/layout.html"
		this.TplName = this.GetTemplate() + "/theme/theme.html"
	}
}

//添加主题
func (this *ThemeController) AddTheme() {
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
	theme_service := theme.NewThemeService()
	_, err = theme_service.AddTheme(attachment_id, title)

	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	this.Ctx.Output.JSON(lib.ApiSuccess(""), false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}

//更新主题
func (this *ThemeController) UpdateTheme() {
	theme_id, err1 := this.GetInt64("Id")
	attachment_id, err2 := this.GetInt64("AttachmentId")
	title := this.GetString("Title")
	if err1 != nil && err2 != nil {
		ers := make(map[string]interface{})
		ers["theme_err"] = err1
		ers["attachment_err"] = err2
		this.Ctx.Output.JSON(lib.ApiErr(ers), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	theme_service := theme.NewThemeService()
	_, err := theme_service.UpdateTheme(theme_id, attachment_id, title)
	if err != nil {
		this.Ctx.Output.JSON(lib.ApiErr(err), false, false)
		this.Ctx.Output.Body([]byte(""))
		return
	}
	this.Ctx.Output.JSON(lib.ApiSuccess(""), false, false)
	this.Ctx.Output.Body([]byte(""))
	return
}
