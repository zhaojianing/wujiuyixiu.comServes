package utils

const (
	RECODE_OK          = "200"
	RECODE_DBERR       = "401"
	RECODE_NODATA      = "402"
	RECODE_DATAEXIST   = "403"
	RECODE_DATAERR     = "404"
	RECODE_SESSIONERR  = "411"
	RECODE_LOGINERR    = "412"
	RECODE_PARAMERR    = "413"
	RECODE_USEERR      = "414"
	RECODE_ROLEERR     = "415"
	RECODE_PWDERR      = "416"
	RECODE_REQERR      = "421"
	RECODE_IPERR       = "422"
	RECOOE_INSERTERROR = "423"
	RECODE_THIRDERR    = "431"
	RECODE_IOERR       = "432"
	RECODE_SERVERERR   = "450"
	RECODE_UNKNOWERR   = "451"
	RECODE_REMOLEERR   = "452"
)

var readText = map[string]string{
	RECODE_OK:          "请求成功",
	RECODE_DBERR:       "数据库查询错误",
	RECODE_NODATA:      "无数据",
	RECODE_DATAEXIST:   "数据已存在",
	RECODE_DATAERR:     "数据错误",
	RECODE_SESSIONERR:  "用户未登录",
	RECODE_LOGINERR:    "用户登陆失败",
	RECODE_PARAMERR:    "参数错误",
	RECODE_USEERR:      "用户不存在或者未激活",
	RECODE_ROLEERR:     "用户身份错误",
	RECODE_PWDERR:      "密码错误",
	RECODE_REQERR:      "非法请求或者请求次数受限",
	RECODE_IPERR:       "IP受限",
	RECOOE_INSERTERROR: "插入数据失败",
	RECODE_THIRDERR:    "第三方系统错误",
	RECODE_IOERR:       "文件读写错误",
	RECODE_SERVERERR:   "内部错误",
	RECODE_UNKNOWERR:   "未知错误",
	RECODE_REMOLEERR:   "删除错误",
}

func RecodeText(code string) string {
	str, ok := readText[code]
	if ok {
		return str
	}
	return readText[RECODE_UNKNOWERR]
}
