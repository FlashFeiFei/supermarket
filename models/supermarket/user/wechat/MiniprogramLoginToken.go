package wechat

import (
	"github.com/astaxie/beego"
	"github.com/flashfeifei/supermarket/lib"
	"github.com/flashfeifei/supermarket/models/supermarket/user/wechat"
	"github.com/robbert229/jwt"
	"reflect"
	"strconv"
	"time"
)

type miniprogramLoginToken struct {
	//用户id
	SupermarketUserId int64
	//微信id
	Openid string
	//微信unionid
	Unionid string
	//账号类型
	AccountType int
	//sessionkey
	SessionKey string
}

func NewMiniprogroamLoginToken(supermarket_user_id int64, openid, unionid, session_key string) *miniprogramLoginToken {
	return &miniprogramLoginToken{
		SupermarketUserId: supermarket_user_id,
		Openid:            openid,
		Unionid:           unionid,
		SessionKey:        session_key,
		AccountType:       wechat.USER_TYPE_MINIPROGRAM,
	}
}

//保存信息到缓存
func (this *miniprogramLoginToken) Save() (error) {
	//缓存
	redis_cache := lib.GetRedisInstance()
	//缓存12小时
	redis_cache.Put(this.GetCacheKey(), this.CreateToken(), 43200*time.Second)
	return nil
}

//获取token
func (this *miniprogramLoginToken) GetToken() (string, error) {
	//缓存
	redis_cache := lib.GetRedisInstance()
	token := redis_cache.Get(this.GetCacheKey())
	return token.(string), nil
}

//获得缓存的key
func (this *miniprogramLoginToken) GetCacheKey() (string) {
	account_type := strconv.FormatInt(int64(this.AccountType), 10)
	return this.Unionid + ":" + account_type + ":" + this.Openid + ":" + this.SessionKey
}

//创建token
func (this *miniprogramLoginToken) CreateToken() (string) {
	key := beego.AppConfig.String("jwt_hsmac256_key")
	//获得加密的key
	algorithm := jwt.HmacSha256(key)
	claims := jwt.NewClaim()
	//通过反射去获取对象的属性和属性值
	getType := reflect.TypeOf(*this)
	beego.Debug("通过反射获取对象的类型")
	beego.Debug(getType)
	getValue := reflect.ValueOf(*this)
	beego.Debug("通过反射获取对象的值")
	beego.Debug(getValue)
	// 获取方法字段
	// 1.先获取interface的reflect.Type,然后通过NumField进行遍历
	// 2.再通过reflect.Type的Field获取其Field
	// 3.最后通过Field的Interface()得到对应的Value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		//设置对象的属性
		claims.Set(field.Name, value)
	}

	token, err := algorithm.Encode(claims)
	if err != nil {
		panic(err)
	}
	return token
}
