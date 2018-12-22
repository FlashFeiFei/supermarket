package util

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/admin"
)

//初始化项目的rbac数据库
func RegisterDBAdminModel() {
	orm.RegisterModel(new(admin.Group), new(admin.Node), new(admin.Role), new(admin.User))
}

//初始化supermarker项目数据库
func RegisterDBSupermarket() {

}

func RegisterDBModel() {
	//注册后台的rbac数据库
	RegisterDBAdminModel()
}
