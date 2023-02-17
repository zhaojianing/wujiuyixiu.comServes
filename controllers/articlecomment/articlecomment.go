package articlecomment

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models/article"
	"servers/models/articlecomment"
	"servers/utils"
	"strconv"
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
	uidName := s.GetString("uid_name")
	uidLike, _ := s.GetInt("uidLike")
	isArticleCommentReply := "0"
	//var reply *models.ArticleCommentReply
	date := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	_, e := articlecomment.InsertArticleComment(articleUid, userId, uidContainer, uidLike, date, isArticleCommentReply, uidName)
	if e != nil {
		reqs["code"] = utils.RECOOE_INSERTERROR
		reqs["msg"] = utils.RecodeText(utils.RECOOE_INSERTERROR)
		reqs["data"] = nil
		return
	}
	// 获取文章信息
	if data, err := article.GetArticle(articleUid); err != nil {
		reqs["code"] = utils.RECODE_DBERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DBERR)
		reqs["data"] = nil
		return
	} else {
		if _, error := article.PutArticleCommentNumber(data.CommentNumber, data.Id); error != nil {
			reqs["watchNum"] = "添加错误"
		}
		reqs["watchNum"] = "评论添加成功"
	}
	reqs["code"] = utils.RECODE_OK
	reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
	reqs["data"] = "success"
}

func (s *ArticleComment) Get() {
	articleUid := s.GetString("article_uid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	// 获取一级评论
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
