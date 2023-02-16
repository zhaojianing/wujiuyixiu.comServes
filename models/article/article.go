package article

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func InsertArticle(articleUid string, articleTitle string, articleContainerHtml string,
	articleContainerText string, articleUserid string, articleCover string, isEncryption string,
	articlePassword string, articleCollect string, like int, watchNumber int, commentNumber int,
	classification string, comment *models.ArticleComment, date string, username string) (id int64, error error) {
	o := orm.NewOrm()
	var article models.Article
	article.ArticleUid = articleUid
	article.ArticleTitle = articleTitle
	article.ArticleContainerHtml = articleContainerHtml
	article.ArticleContainerText = articleContainerText
	article.ArticleUserid = articleUserid
	article.ArticleCover = articleCover
	article.IsEncryption = isEncryption
	article.ArticlePassword = articlePassword
	article.ArticleCollect = articleCollect
	article.Like = like
	article.WatchNumber = watchNumber
	article.CommentNumber = commentNumber
	article.Classification = classification
	article.Comment = comment
	article.Date = date
	article.ArticleUseridName = username

	id, error = o.Insert(&article)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	return id, error
}

func GetArticle(articleUid string) (data models.Article, err error) {
	o := orm.NewOrm()
	var article models.Article
	article.ArticleUid = articleUid

	_, err = o.QueryTable(&article).Filter("ArticleUid", articleUid).All(&article)
	if err != nil {
		return article, err
	}
	return article, nil
}

func PutArticleWatchNum(num int, id int) (rid int, err error) {
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	article.WatchNumber = num + 1

	println("num + 1", num+1)
	if num, err := o.Update(&article, "WatchNumber"); err == nil {
		return int(num), nil
	}
	//}
	return 0, err
}
