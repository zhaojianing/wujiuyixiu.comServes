package userinfo

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models/userinfo"
	"servers/utils"
)

type InfoController struct {
	beego.Controller
}

func (s *InfoController) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

func (s *InfoController) Get() {
	uid := s.GetString("uid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	if data, err := userinfo.GetUser(uid); err != nil {
		reqs["code"] = utils.RECODE_USEERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_USEERR)
		reqs["data"] = nil
		return
	} else {
		// 清除密码
		data.Password = ""
		reqs["code"] = utils.RECODE_OK
		reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
		reqs["data"] = data
	}

}
