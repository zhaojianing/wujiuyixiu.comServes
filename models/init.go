package models

import (
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Statistics struct {
	Id          int `orm:"pk;auto"`
	Totalpeople int
}

type Users struct {
	Id         int       `json:"id"`
	Uid        string    `json:"uid"`
	Name       string    `json:"name"`                   // 用户名称
	Account    string    `json:"account"`                // 账号
	Password   string    `json:"password"`               // 密码
	Signature  string    `json:"signature"`              // 个性签名
	AvatarUrl  string    `json:"avatar_url"`             // 头像地址
	Email      string    `json:"email"`                  // 用户邮箱
	ArticleNum int       `json:"article_num"`            // 文章个数
	StartNum   int       `json:"start_num"`              // 关注个数
	FanNum     int       `json:"fan_num"`                // 粉丝个数
	DynamicNum int       `json:"dynamic_num"`            // 动态个数
	MartialArt string    `json:"martial_art"`            // 门派
	Date       time.Time `orm:"auto_now_add;type(date)"` // 用户创建时间
}

type MaoShanUser struct { // 保存茅山排盘
	Id           int       `json:"id"`                     // 储存id
	Uid          string    `json:"uid"`                    // 所属用户id
	Time         string    `json:"time"`                   // 保存时间
	Title        string    `json:"title"`                  // 用户保存茅山奇门标题
	Introduction string    `json:"introduction"`           // 用户保存解卦内容
	IsPublic     int       `json:"is_public"`              // 是否公开0/1
	Date         time.Time `orm:"auto_now_add;type(date)"` // 用户创建时间
}

type XuanKongUser struct { // 保存玄空排盘
	Id           int       `json:"id"`                     // 储存id
	Uid          string    `json:"uid"`                    // 所属用户id
	Time         string    `json:"time"`                   // 保存时间
	Title        string    `json:"title"`                  // 用户保存玄空标题
	Introduction string    `json:"introduction"`           // 用户保存解卦内容
	IsPublic     int       `json:"is_public"`              // 是否公开0/1
	Date         time.Time `orm:"auto_now_add;type(date)"` // 用户创建时间
}

type MessageForMe struct { // 保存玄空排盘
	Id           int    `json:"id"`           // 储存id
	Uid          string `json:"uid"`          // 所属用户id
	Name         string `json:"name"`         // 用户保存茅山奇门标题
	Introduction string `json:"introduction"` // 用户保存解卦内容
	City         string `json:"city"`         // 存储城市
	Date         string `json:"date"`         // 存储时间
}

type Article struct { // 文章表
	Id                   int             `json:"id"`                                      // 文章id
	ArticleUid           string          `json:"article_uid"`                             // 文章uid
	ArticleTitle         string          `json:"article_title"`                           // 文章标题
	ArticleContainerHtml string          `orm:"type(text)" json:"article_container_html"` // 文章内容html
	ArticleContainerText string          `orm:"type(text)" json:"article_container_text"` // 文章内容text
	ArticleUserid        string          `json:"article_userid"`                          // 文章所属用户 uid
	ArticleUseridName    string          `json:"article_userid_name"`                     // 文章所属用户名字
	ArticleCover         string          `json:"article_cover"`                           // 文章封面地址
	IsEncryption         string          `json:"is_encryption"`                           // 文章是否加密
	ArticlePassword      string          `json:"article_password"`                        // 文章密码
	ArticleCollect       string          `json:"article_collect"`                         // 文章收藏（1，0）
	Like                 int             `json:"like"`                                    // 文章点赞
	WatchNumber          int             `json:"watch_number"`                            // 文章查看人数
	CommentNumber        int             `json:"comment_number"`                          // 评论总数
	Classification       string          `json:"classification"`                          // 所属分类
	Comment              *ArticleComment `orm:"null;rel(fk)"`                             // 设置一对多关系
	Date                 string          `json:"date"`                                    // 用户创建时间
}

type ArticleComment struct { // 一级评论表
	Id                    int                  `json:"id"`                       // 评标表id
	ArticleUid            string               `json:"article_uid"`              // 所属文章uid
	Uid                   string               `json:"uid"`                      // 评论人uid
	UidContainer          string               `json:"uid_container"`            // 评论人内容
	UidLike               int                  `json:"uid_like"`                 // 评论人被点赞数
	IsArticleCommentReply string               `json:"is_article_comment_reply"` // 是否有二级回复评论
	Reply                 *ArticleCommentReply `orm:"rel(fk)"`                   // 设置一对多关系(被回复的评论)
	Date                  int                  `json:"date"`                     // 用户创建时间
}

type ArticleCommentReply struct { // 二级回复评论表
	Id               int    `json:"id"`                 // 评标表id
	ArticleUid       string `json:"article_uid"`        // 所属文章uid
	ArticleCommentId int    `json:"article_comment_id"` // 所属一级评论id
	ReplyUserUid     string `json:"reply-user-uid"`     // 被回复的评论人uid
	Uid              string `json:"uid"`                // 评论人uid
	UidContainer     string `json:"uid_container"`      // 评论人内容
	UidLike          int    `json:"uid_like"`           // 评论人被点赞数
	Date             int    `json:"date"`               // 用户创建时间
}

func RegisterDB() {
	host, _ := beego.AppConfig.String("mysqlurls")
	port, _ := beego.AppConfig.String("mysqlport")
	dbname, _ := beego.AppConfig.String("mysqldb")
	user, _ := beego.AppConfig.String("mysqluser")
	pwd, _ := beego.AppConfig.String("mysqlpass")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbcon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dbcon)
	conn, err := sql.Open("mysql", dbcon)
	if err != nil {
		fmt.Printf("连接数据库失败，error:%s", err)
		return
	} else {
		fmt.Println("连接数据库成功")
		fmt.Println("更新啦！！！！！！！！！！更新必须把更新的表写在最前面，否则不更新")
		orm.RegisterModel(new(Article), new(ArticleComment), new(ArticleCommentReply), new(XuanKongUser),
			new(MessageForMe), new(Users), new(MaoShanUser), new(Statistics))
	}
	//关闭数据库
	defer conn.Close()

}
