package lib

//成功的返回数据格式
func ApiSuccess(result interface{}) (success map[string]interface{}) {
	success["code"] = 0
	success["data"] = result
	return success
}

//错误的返回数据格式
func ApiErr(result interface{}) (success map[string]interface{}) {
	success["code"] = 0
	success["data"] = result
	return success
}
