package lib

//成功的返回数据格式
func ApiSuccess(data interface{}) (interface{}) {
	success := make(map[string]interface{})
	success["code"] = 0
	success["data"] = data
	return success
}

//错误的返回数据格式,改api用于自己服务器发成错误时候的
func ApiErr(data interface{}) (interface{}) {
	err := make(map[string]interface{})
	err["code"] = 1
	err["data"] = data
	return err
}

//这个函数用户，服务端发送请求，发生错误时候返回的api
func ApiErrOpenPlatform(code int, data interface{}) (interface{}) {
	err := make(map[string]interface{})
	err["code"] = code
	err["data"] = data
	return err
}
