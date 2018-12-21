/*
	Copyright 2017 by rabbit author: gdccmcm14@live.com.
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License
*/
package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/lib"
)

func InitData() {
	InsertUser()
	InsertGroup()
	InsertRole()
	InsertNodes()
}

// 用户数据
func InsertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.Username = beego.AppConfig.String("rbac_admin_user")
	u.Nickname = "Liangyu"
	u.Password = lib.Pwdhash(beego.AppConfig.String("rbac_admin_user"))
	u.Email = "51785816@qq.com"
	u.Remark = "代码搬运工.."
	// 2 stand for close, but it has very high authority
	u.Status = 2
	u.Createtime = lib.GetTime()
	err := u.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	u1 := new(User)
	u1.Username = "test"
	u1.Nickname = "TuziTest"
	u1.Password = lib.Pwdhash("test")
	u1.Email = "569929309@qq.com"
	u1.Remark = "Just a Test User"
	u1.Status = 1
	u1.Createtime = lib.GetTime()
	err1 := u1.Insert()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert user end")
}

// 模组数据
func InsertGroup() {
	fmt.Println("insert group ...")
	g := new(Group)
	g.Name = "兔子后台"
	g.Title = "后台管理"
	g.Sort = 1
	g.Id = 1
	g.Status = 1
	e := g.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println("insert group end")
}

// 角色数据
func InsertRole() {
	fmt.Println("insert role ...")
	r := new(Role)
	r.Name = "管理员"
	r.Remark = "权限最高的一群人"
	r.Status = 1
	r.Title = "管理员角色"
	r.Insert()
	fmt.Println("insert role end")
}

// 节点数据
func InsertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	nodes := []Node{
		/*

			RBAC管理中心

		*/
		{Id: 1, Name: "rbac", Title: "权限中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		{Id: 2, Name: "node/index", Title: "节点管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 3, Name: "Index", Title: "节点首页", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 4, Name: "AddAndEdit", Title: "增编节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 5, Name: "DelNode", Title: "删除节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},

		{Id: 6, Name: "user/index", Title: "用户管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 7, Name: "Index", Title: "用户首页", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 8, Name: "AddUser", Title: "增加用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 9, Name: "UpdateUser", Title: "更新用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 10, Name: "DelUser", Title: "删除用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},

		{Id: 11, Name: "group/index", Title: "分组管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 12, Name: "Index", Title: "分组首页", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 13, Name: "AddGroup", Title: "增加分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 14, Name: "UpdateGroup", Title: "更新分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 15, Name: "DelGroup", Title: "删除分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},

		{Id: 16, Name: "role/index", Title: "角色管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 17, Name: "index", Title: "角色首页", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 18, Name: "AddAndEdit", Title: "增编角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 19, Name: "DelRole", Title: "删除角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 20, Name: "GetList", Title: "列出角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 21, Name: "AccessToNode", Title: "显示权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 22, Name: "AddAccess", Title: "增加权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 23, Name: "RoleToUserList", Title: "列出角色下用户", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 24, Name: "AddRoleToUser", Title: "授予用户角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
	}
	for _, v := range nodes {
		n := new(Node)
		n.Id = v.Id // 这句是无效的,后来 bug 被 beego 官方改好了
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		e := n.Insert()
		if e != nil {
			fmt.Printf("%#v:%#v\n", n, e.Error())
		}
	}
	fmt.Println("insert node end")
}
