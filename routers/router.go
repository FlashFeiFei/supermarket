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
// 路由包
package routers

import (
	"github.com/flashfeifei/supermarket/routers/admin/attachment"
	"github.com/flashfeifei/supermarket/routers/admin/banner"
	"github.com/flashfeifei/supermarket/routers/supermarket/user/wechat"
)

// 路由
func Run() {
	//后台模块
	rbacrouter()
	//附件管理路由注册
	attachment.AdminAttachmentRouterRegister()
	//banner管理
	banner.AdminBannerRouterRegister()


	//小程序模块
	wechat.MiniprogramLoginRouter()
}
