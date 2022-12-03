package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models"
	"servers/utils"
)

type StatisticsController struct {
	beego.Controller
}

func (s *StatisticsController) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	s.ServeJSON()
}

func (s *StatisticsController) Get() {
	//id, _ := s.GetInt("id")
	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)
	num, e := models.UpdateStatistics()
	if e != nil {
		reqs["code"] = utils.RECOOE_INSERTERROR
		reqs["msg"] = utils.RecodeText(utils.RECOOE_INSERTERROR)
		reqs["data"] = e
	} else {
		reqs["code"] = utils.RECODE_OK
		reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
		reqs["data"] = num
	}
	//s.ServeJSON()
}

func (s *StatisticsController) Post() {
	name := s.GetString("name")
	s.Data["json"] = name
	s.ServeJSON()
}
