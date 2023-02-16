package articlecomment

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models"
	"servers/models/articlecomment"
	"servers/utils"
	"time"
)

type ArticleComment struct {
	beego.Controller
}

func (s *ArticleComment) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

func (s *ArticleComment) Post() {
	articleUid := s.GetString("articleUid")
	userId := s.GetString("uid")
	uidContainer := s.GetString("uidContainer")
	uidLike, _ := s.GetInt("uidLike")
	isArticleCommentReply := "0"
	var reply *models.ArticleCommentReply
	date := time.Now().Unix()

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	_, e := articlecomment.InsertArticleComment(articleUid, userId, uidContainer, uidLike, reply, int(date), isArticleCommentReply)
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

func (s *ArticleComment) Get() {
	articleUid := s.GetString("articleUid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	data, e := articlecomment.GetArticleComment(articleUid)
	if e != nil {
		reqs["code"] = utils.RECODE_DATAERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DATAERR)
		reqs["data"] = nil
		return
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = data
}
