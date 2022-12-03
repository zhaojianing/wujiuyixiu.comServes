package xuankonguser

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models/xuankonguser"
	"servers/utils"
)

type XuanKongUsersController struct {
	beego.Controller
}

func (s *XuanKongUsersController) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

func (s *XuanKongUsersController) Post() {
	uid := s.GetString("uid")
	title := s.GetString("title")
	time := s.GetString("time")
	introduction := s.GetString("introduction")
	isPublic := 0

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	_, e := xuankonguser.InsertXuanKonguser(uid, time, title, introduction, isPublic)
	if e != nil {
		reqs["code"] = utils.RECODE_DATAERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DATAERR)
		reqs["data"] = nil
		return
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = "success"
}

func (s *XuanKongUsersController) Get() {
	uid := s.GetString("uid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	data, e := xuankonguser.GetXuanKonguser(uid)
	if e != nil {
		reqs["code"] = utils.RECODE_DBERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DBERR)
		reqs["data"] = nil
		return
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = data
}
