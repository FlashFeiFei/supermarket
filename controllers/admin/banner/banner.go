package banner

type BannerController struct {
	baseController
}

func (this *BannerController) Index() {
	if this.IsAjax() {

	} else {
		this.Layout = this.GetTemplate() + "/banner/layout.html"
		this.TplName = this.GetTemplate() + "/banner/banner.html"
	}
}
func (this *BannerController) AddBanner() {

}
