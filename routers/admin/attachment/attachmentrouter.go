package attachment

import (
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/controllers/admin/attachment"
)

func AdminAttachmentRouterRegister(){
	beego.Router("/attachment/attachment/Index", &attachment.AttachmentController{}, "*:Index")
	//上传一张图片
	beego.Router("/attachment/attachment/UploadImage", &attachment.AttachmentController{}, "*:UploadImage")
	//更新图片信息
	beego.Router("/attachment/attachment/UpdateImageInfo", &attachment.AttachmentController{}, "*:UpdateImageInfo")
}
