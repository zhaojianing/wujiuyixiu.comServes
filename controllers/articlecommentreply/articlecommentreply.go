package articlecommentreply

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models/article"
	articlecomment "servers/models/articlecomment"
	"servers/models/articlecommentreply"
	"servers/utils"
	"strconv"
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
	replyUserUidName := s.GetString("reply_user_uid_name")
	uidLike, _ := s.GetInt("uidLike")
	uidName := s.GetString("uid_name")
	date := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	id, _ := s.GetInt("articleCommentId")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	// 写入
	_, e := articlecommentreply.InsertArticleCommentReply(articleUid, replyUserUid, uid, uidContainer, uidLike, date, uidName, id, replyUserUidName)
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

func (s *ArticleCommentReply) Get() {
	id, _ := s.GetInt("articleCommentId")
	articleUid := s.GetString("articleUid")
	//uid := s.GetString("uid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	println("id is ", id)
	println("articleuid is ", articleUid)
	data, e := articlecommentreply.GetArticleCommentReply(id, articleUid)
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
