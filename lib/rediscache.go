package lib

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

//获取redis实例
func GetRedisInstance(collectionName, conn, dbnumber, password string) (redis_cache cache.Cache, err error) {
	//初始化连接
	redis_cache, err = cache.NewCache("redis", `{"key":"`+collectionName+`","conn":"`+conn+`", "password":"`+password+`", "dbNum":"`+dbnumber+`"
}`)
	if err != nil {
		return nil, err
	}
	return redis_cache, nil
}
