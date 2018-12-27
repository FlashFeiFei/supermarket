package lib

import (
	"github.com/astaxie/beego"
	"github.com/edwingeng/wuid/redis"
)

var wuid_by_redis *wuid.WUID

func init() {
	registerWuidByRedis()
	beego.Debug("----------------查看一下生成的数据")
	beego.Debug(GetWuidByRedis())
}

//实例化reids的wuid
func registerWuidByRedis() {
	wuid_by_redis = wuid.NewWUID("default", nil)
	err := wuid_by_redis.LoadH24FromRedis(beego.AppConfig.String("cache::redis_host"), beego.AppConfig.String("cache::redis_password"), "wuid")
	if err != nil {
		panic(err)
	}
}

//库的github下的，上面说是线程安全的
//得到一个id号
func GetWuidByRedis() (uint64) {
	return wuid_by_redis.Next()
}
