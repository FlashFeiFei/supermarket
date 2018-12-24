package wechat


type code2Session struct {
	//用户id
	supermarket_user_id int64
	//微信id
	openid       string
	//微信unionid
	unionid      string
	//账号类型
	account_type int
	//sessionkey
	session_key  string
}

func (this *code2Session) save() {

}
