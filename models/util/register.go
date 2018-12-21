package util

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/admin"
)

func RegisterDBAdminModel() {
	orm.RegisterModel(new(admin.Group), new(admin.Node), new(admin.Role), new(admin.User))
}

func RegisterDBModel() {
	RegisterDBAdminModel()
}
