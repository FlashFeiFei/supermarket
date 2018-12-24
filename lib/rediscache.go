package lib

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"strings"
)
//注册redis驱动
import _ "github.com/astaxie/beego/cache/redis"

var redisInstance map[string]cache.Cache

func init() {
	host := beego.AppConfig.String("cache::redis_host")
	pwd := beego.AppConfig.String("cache::redis_password")
	dbnumber := beego.AppConfig.String("cache::redis_dbnumber")
	//注册默认的redis连接池
	_, err := GetRedisInstance(host, dbnumber, pwd)
	if err != nil {
		panic(err)
	}
}

//获取redis实例
func GetRedisInstance(conn, dbnumber, password string) (redis_cache cache.Cache, err error) {
	var content_info []string
	content_info = append(content_info, conn)
	content_info = append(content_info, dbnumber)
	content_info = append(content_info, password)
	key := Md5(strings.Join(content_info, " "))
	if redis_cache, exist := redisInstance[key]; exist {
		return redis_cache, nil;
	}
	//初始化连接
	redis_cache, err = cache.NewCache("redis", `{"conn":"`+conn+`","password":"`+password+`","dbNum":"`+dbnumber+`"}`)
	if err != nil {
		return nil, err
	}
	redisInstance[key] = redis_cache
	return redis_cache, nil
}
