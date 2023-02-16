package articlerecommendation

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models/articlerecommendation"
	"servers/utils"
)

type ArticleRecommendation struct {
	beego.Controller
}

func (s *ArticleRecommendation) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

func (s *ArticleRecommendation) Get() {
	pagination, _ := s.GetInt("pagination")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	// 当前分页数据
	data, e := articlerecommendation.GetArticleRecommendation(pagination)
	if e != nil {
		reqs["code"] = utils.RECODE_DATAERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DATAERR)
		reqs["data"] = nil
		return
	}
	// 分页所需要的总数
	cnt, err := articlerecommendation.GetArticleRecommendationCount()
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
