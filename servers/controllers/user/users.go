package user

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models/user"
	"servers/utils"
)

type UsersController struct {
	beego.Controller
}

func (s *UsersController) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

// Get 获取用户信息(登录)
func (s *UsersController) Get() {
	account := s.GetString("account")
	password := s.GetString("password")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	data, e := user.SignInUserModelData(account, password)
	if e != nil {
		reqs["code"] = utils.RECODE_DBERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DBERR)
		reqs["data"] = nil
		return
	}

	if password != data[0].Password {
		reqs["code"] = utils.RECODE_PWDERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_PWDERR)
		reqs["data"] = nil
		return
	}

	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = data[0]
}

func (s *UsersController) Post() {
	account := s.GetString("account")
	password := s.GetString("password")
	name := s.GetString("username")
	uid := s.GetString("uid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	// 账号去重
	flag, _ := user.DeduplicationUser(account)
	if flag {
		reqs["code"] = utils.RECODE_USEERR
		reqs["msg"] = "账号已存在！"
		reqs["data"] = ""
		return
	}

	// 注册账号
	userId, e := user.InsertUserModelData(account, password, name, uid)
	if e != nil {
		reqs["code"] = utils.RECODE_USEERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_USEERR)
		reqs["data"] = e
		return
	}

	// 注册成功，并登录
	data, e := user.QueryUserModelData(userId)
	if e != nil {
		reqs["code"] = utils.RECODE_DBERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DBERR)
		reqs["data"] = e
		return
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = data[0]
}
