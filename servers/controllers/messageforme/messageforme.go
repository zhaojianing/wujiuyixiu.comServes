package messageforme

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"servers/models/messageforme"
	"servers/utils"
)

type MessageForMeController struct {
	beego.Controller
}

func (s *MessageForMeController) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

func (s *MessageForMeController) Post() {
	uid := s.GetString("uid")
	name := s.GetString("name")
	introduction := s.GetString("introduction")
	date := s.GetString("date")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	_, e := messageforme.InsertMessageForMe(uid, name, introduction, date)
	if e != nil {
		reqs["code"] = utils.RECOOE_INSERTERROR
		reqs["msg"] = utils.RecodeText(utils.RECOOE_INSERTERROR)
		reqs["data"] = nil
		return
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = "success"
}

func (s *MessageForMeController) Get() {
	pagination, _ := s.GetInt("pagination")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	fmt.Println("pagination is %i", pagination)

	data, e := messageforme.GetMessageForMe(pagination)
	if e != nil {
		reqs["code"] = utils.RECODE_DATAERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DATAERR)
		reqs["data"] = nil
		return
	}
	cnt, err := messageforme.GetMessageForMeCount()
	if err != nil {
		reqs["code"] = utils.RECODE_DATAERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DATAERR)
		reqs["data"] = nil
		return
	}

	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = data
	reqs["status"] = cnt
}
