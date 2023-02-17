package articlecommentreply

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func InsertArticleCommentReply(articleUid string, replyUserUid string, uid string, uidcontainer string, uidlike int,
	date string, uidName string, acid int, replyUserUidName string) (id int64, error error) {
	o := orm.NewOrm()
	var articleCommentReply models.ArticleCommentReply
	articleCommentReply.ArticleUid = articleUid
	articleCommentReply.ReplyUserUid = replyUserUid
	articleCommentReply.Uid = uid
	articleCommentReply.UidContainer = uidcontainer
	articleCommentReply.UidLike = uidlike
	articleCommentReply.Date = date
	articleCommentReply.UidName = uidName
	articleCommentReply.ArticleCommentId = acid
	articleCommentReply.ReplyUserUidName = replyUserUidName

	id, error = o.Insert(&articleCommentReply)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	return id, error
}

func GetArticleCommentReply(articleCommentId int, articleUid string) (data []*models.ArticleCommentReply, error error) {
	o := orm.NewOrm()
	var articleCommentReply []*models.ArticleCommentReply
	_, err := o.QueryTable("ArticleCommentReply").Filter("ArticleCommentId", articleCommentId).Filter(
		"ArticleUid", articleUid).All(&articleCommentReply)
	// .Filter("ReplyUserUid", uid)
	if err != nil {
		return nil, err
	}
	return articleCommentReply, nil

}
