package wechat

import (
	"github.com/astaxie/beego"
	"github.com/robbert229/jwt"
	"reflect"
	"strconv"
)

type token struct {
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

//保存信息到缓存
func (this *token) Save() {
	
}

//获得缓存的key
func (this *token) GetCacheKey() (string) {
	account_type := strconv.FormatInt(int64(this.AccountType), 10)
	return this.Unionid + ":" + account_type + ":" + this.Openid + ":" + this.SessionKey
}

//创建token
func (this *token) CreateToken() (string) {
	key := beego.AppConfig.String("jwt_hsmac256_key")
	//获得加密的key
	algorithm := jwt.HmacSha256(key)
	claims := jwt.NewClaim()
	//通过反射去获取对象的属性和属性值
	getType := reflect.TypeOf(this)
	beego.Debug("通过反射获取对象的类型")
	beego.Debug(getType)
	getValue := reflect.ValueOf(this)
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
