package attachment

type AttachmentController struct {
	baseController
}

func (this *AttachmentController) Index() {
	if this.IsAjax() {

	} else {
		this.Layout = this.GetTemplate() + "/attachment/layout.html"
		this.TplName = this.GetTemplate() + "/attachment/user.html"
	}
}
