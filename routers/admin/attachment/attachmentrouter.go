package attachment

import (
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/controllers/admin/attachment"
)

func AdminAttachmentRouterRegister(){
	beego.Router("/attachment/attachment/index", &attachment.AttachmentController{}, "*:Index")
}
