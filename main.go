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

// Main Web Entrance
package main

import (
	"flag"
	"mime"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/flashfeifei/supermarket/conf"
	"github.com/flashfeifei/supermarket/controllers"
	"github.com/flashfeifei/supermarket/lib"
	"github.com/flashfeifei/supermarket/models"
	"github.com/flashfeifei/supermarket/routers"
)

func init() {
	// init flag
	//通过命令行传入的参数来覆盖一些默认的参数，从而执行一些非默认的逻辑
	flags := conf.FlagConfig{}

	// user that hide
	flags.User = flag.String("user", "", "user")

	// db init or rebuild
	// 初始化数据库
	flags.DbInit = flag.Bool("db", false, "初始化数据库")
	flags.DbInitForce = flag.Bool("f", false, "强制init db先删除db，然后重新构建它")

	// rbac config rebuild
	// rbac配置重建
	flags.Rbac = flag.Bool("rbac", false, "rbac配置重建")

	// front-end  view
	home := flag.String("home", "", "home template")

	// config file position
	//如果配置文件位置为空，则使用默认值
	config := flag.String("config", "", "如果配置文件位置为空，则使用默认值")

	// 获取命令传入的参数
	flag.Parse()

	// init config
	if *config != "" {
		beego.Trace("use diy config")
		err := beego.LoadAppConfig("ini", *config)
		if err != nil {
			beego.Trace(err.Error())
		} else {
			beego.Trace("Use config:" + *config)
		}
	}

	if *home != "" {
		beego.Trace("Home template is " + *home)
		beego.AppConfig.Set(beego.BConfig.RunMode+"::"+"home_template", *home)
	}

	conf.InitConfig()

	// init lang
	// just add some ini in conf such locale_zh-CN.ini and edit app.conf
	// beego自带的语言包
	langTypes := strings.Split(beego.AppConfig.String("lang_types"), "|")

	for _, lang := range langTypes {
		beego.Trace("Load language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Load language error:", err)
			return
		}
	}

	// add func map
	// 注册模板可以调用的函数
	beego.Trace("add i18n function map")
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Trace("add stringsToJson function  map")
	beego.AddFuncMap("stringsToJson", lib.StringsToJson)
	mime.AddExtensionType(".css", "text/css") // some not important

	// init model
	//模型运行
	beego.Trace("model run")
	models.Run(flags)

	// init router
	beego.Trace("router run")
	routers.Run()

	beego.Trace("start open error template")
	beego.ErrorController(&controllers.ErrorController{})
}

// Start!
func main() {
	beego.Trace("Start Listen ...")
	beego.Run()
}
