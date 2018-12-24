package wechat

type code2Session struct {
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

func (this *code2Session) save() {

}
