package articlecommentreply

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func InsertArticleCommentReply(articleUid string, replyUserUid string, uid string, uidcontainer string, uidlike int,
	date int) (id int64, error error) {
	o := orm.NewOrm()
	var articleCommentReply models.ArticleCommentReply
	articleCommentReply.ArticleUid = articleUid
	articleCommentReply.ReplyUserUid = replyUserUid
	articleCommentReply.Uid = uid
	articleCommentReply.UidContainer = uidcontainer
	articleCommentReply.UidLike = uidlike
	articleCommentReply.Date = date

	id, error = o.Insert(&articleCommentReply)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	return id, error
}

func GetArticleCommentReply(articleCommentId int, articleUid string, uid string) (data []*models.ArticleCommentReply, error error) {
	o := orm.NewOrm()
	var articleCommentReply []*models.ArticleCommentReply
	_, err := o.QueryTable(&articleCommentReply).Filter("ArticleUid", articleUid).Filter("ArticleCommentId",
		articleCommentId).Filter("Uid", uid).OrderBy("-uid_like").All(&articleCommentReply)
	if err != nil {
		return nil, err
	}
	return articleCommentReply, nil

}
