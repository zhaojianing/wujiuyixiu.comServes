package article

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/models"
	article "servers/models/article"
	"servers/utils"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (s *ArticleController) DataJson(reqs map[string]interface{}) {
	s.Data["json"] = reqs
	err := s.ServeJSON()
	if err != nil {
		return
	}
}
func (s *ArticleController) Post() {
	articleUid, _ := utils.RadomTimeNum()
	articleTitle := s.GetString("title")
	articleContainerHtml := s.GetString("html")
	articleContainerText := s.GetString("text")
	articleUserid := s.GetString("article_user_id")
	articleCover := s.GetString("article_cover")
	isEncryption := s.GetString("is_encryption")
	articlePassword := s.GetString("article_password")
	articleCollect := s.GetString("article_collect")
	like, _ := s.GetInt("like")
	watchNumber, _ := s.GetInt("watch_number")
	commentNumber, _ := s.GetInt("comment_number")
	classification := s.GetString("classification")
	username := s.GetString("username")
	var comment *models.ArticleComment
	// 获取时间戳，前台转换
	//date := time.Now().Unix()
	date := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	_, e := article.InsertArticle(articleUid, articleTitle, articleContainerHtml,
		articleContainerText, articleUserid, articleCover, isEncryption,
		articlePassword, articleCollect, like, watchNumber, commentNumber,
		classification, comment, date, username)
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

func (s *ArticleController) Get() {
	articleUid := s.GetString("article_uid")

	reqs := make(map[string]interface{})
	defer s.DataJson(reqs)

	// 获取文章信息
	if data, err := article.GetArticle(articleUid); err != nil {
		reqs["code"] = utils.RECODE_DBERR
		reqs["msg"] = utils.RecodeText(utils.RECODE_DBERR)
		reqs["data"] = nil
		return
	} else {
		if _, error := article.PutArticleWatchNum(data.WatchNumber, data.Id); error != nil {
			reqs["watchNum"] = "添加错误"
		}
		reqs["code"] = utils.RECODE_OK
		reqs["msg"] = utils.RecodeText(utils.RECODE_OK)
		reqs["data"] = data
		reqs["watchNum"] = "观看添加成功"
		println("dataWatchNum is", data.WatchNumber)
	}

}
