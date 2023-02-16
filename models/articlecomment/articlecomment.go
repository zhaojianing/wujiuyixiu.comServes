package articlecomment

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func InsertArticleComment(articleUid string, uid string, uidcontainer string, uidlike int,
	reply *models.ArticleCommentReply, date int, isArticleCommentReply string) (id int64, error error) {
	o := orm.NewOrm()
	var articleComment models.ArticleComment
	articleComment.ArticleUid = articleUid
	articleComment.Uid = uid
	articleComment.UidContainer = uidcontainer
	articleComment.UidLike = uidlike
	articleComment.Reply = reply
	articleComment.Date = date
	articleComment.IsArticleCommentReply = isArticleCommentReply

	id, error = o.Insert(&articleComment)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	return id, error
}

func GetArticleComment(articleUid string) (data []*models.ArticleComment, error error) {
	o := orm.NewOrm()
	var articleComment []*models.ArticleComment
	_, err := o.QueryTable("ArticleComment").OrderBy("-date", "-id").All(&articleComment)
	if err != nil {
		return nil, err
	}
	return articleComment, nil
}

// ChangeIsArticleCommentReply 修改是否评论字段
func ChangeIsArticleCommentReply(id int) (rid int, err error) {
	o := orm.NewOrm()
	articleComment := models.ArticleComment{Id: id}
	//if o.Read(&articleComment) == nil {
	articleComment.IsArticleCommentReply = "1"
	// 只更新IsArticleCommentReply字段
	if num, err := o.Update(&articleComment, "IsArticleCommentReply"); err == nil {
		return int(num), nil
	}
	//}
	return 0, err
}
