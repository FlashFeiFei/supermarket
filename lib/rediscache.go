package lib

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var redis_cache cache.Cache

func init() {
	registerRedisInstance()
}

//获取redis实例
func registerRedisInstance() () {
	conn := beego.AppConfig.String("cache::redis_host")
	pwd := beego.AppConfig.String("cache::redis_password")
	dbnumber := beego.AppConfig.String("cache::redis_dbnumber")
	collectionName := beego.AppConfig.String("cache::collection_name")
	//初始化连接
	adapter, err := cache.NewCache("redis", `{"key":"`+collectionName+`","conn":"`+conn+`", "password":"`+pwd+`", "dbNum":"`+dbnumber+`"
}`)
	if err != nil {
		panic(err)
	}
	redis_cache = adapter

}

//获得实例
func GetRedisInstance() (cache.Cache) {
	return redis_cache
}
