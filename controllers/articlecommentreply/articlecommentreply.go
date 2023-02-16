package articlecommentreply

import (
	beego "github.com/beego/beego/v2/server/web"
	articlecomment "servers/models/articlecomment"
	"servers/models/articlecommentreply"
	"servers/utils"
	"time"
)

type ArticleCommentReply struct {
	beego.Controller
}

func (s *ArticleCommentReply) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}

func (s *ArticleCommentReply) Post() {
	articleUid := s.GetString("articleUid")
	replyUserUid := s.GetString("replyUserUid")
	uid := s.GetString("uid")
	uidContainer := s.GetString("uidContainer")
	uidLike, _ := s.GetInt("uidLike")
	date := time.Now().Unix()

	id, _ := s.GetInt("articleCommentId")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	// 写入
	_, e := articlecommentreply.InsertArticleCommentReply(articleUid, replyUserUid, uid, uidContainer, uidLike, int(date))
	if e != nil {
		reqs["code"] = utils.RECOOE_INSERTERROR
		reqs["msg"] = utils.RecodeText(utils.RECOOE_INSERTERROR)
		reqs["data"] = nil
		return
	}
	// 修改以及评论表-是否评论字段
	if id, err := articlecomment.ChangeIsArticleCommentReply(id); err != nil {
		reqs["code"] = utils.RECODE_IOERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_IOERR)
		reqs["data"] = id
		return
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = "success"
}

func (s *ArticleCommentReply) Get() {
	id, _ := s.GetInt("articleCommentId")
	articleUid := s.GetString("articleUid")
	uid := s.GetString("uid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	data, e := articlecommentreply.GetArticleCommentReply(id, articleUid, uid)
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
