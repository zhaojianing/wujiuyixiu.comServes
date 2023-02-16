package articlerecommendation

import (
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func GetArticleRecommendation(pagination int) (data []*models.Article, error error) {
	o := orm.NewOrm()
	var dataArticle []*models.Article
	_, err := o.QueryTable("Article").OrderBy("-date", "-id").Limit(10, (pagination-1)*10).All(&dataArticle)
	if err != nil {
		return nil, err
	}
	return dataArticle, nil
}

// GetArticleRecommendationCount 分页需要的总数
func GetArticleRecommendationCount() (count int64, error error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Article").Count() // SELECT COUNT(*) FROM USER
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
