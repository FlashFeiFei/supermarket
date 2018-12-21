package util

import (
	"github.com/astaxie/beego/orm"
	"github.com/flashfeifei/supermarket/models/admin"
	"github.com/flashfeifei/supermarket/models/blog"
)

func RegisterDBAdminModel() {
	orm.RegisterModel(new(admin.Group), new(admin.Node), new(admin.Role), new(admin.User))
}

func RegisterDBBlogModel() {
	orm.RegisterModel(new(blog.Category), new(blog.Config), new(blog.Paper), new(blog.Roll))
}

func RegisterDBModel() {
	RegisterDBAdminModel()
	RegisterDBBlogModel()
}
