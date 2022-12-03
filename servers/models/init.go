package models

import (
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
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
	Id               int             `json:"id"`                                 // 文章id
	ArticleTitle     string          `json:"article_title"`                      // 文章标题
	ArticleContainer string          `orm:"type(text)" json:"article_container"` // 文章内容
	ArticleUid       string          `json:"article_uid"`                        // 文章所属用户 uid
	Like             int             `json:"like"`                               // 文章点赞
	Comment          *ArticleComment `orm:"rel(fk)"`                             //设置一对多关系
	Date             time.Time       `orm:"auto_now_add;type(date)"`             // 用户创建时间
}

type ArticleComment struct { // 评论表
	Id           int                  `json:"id"`                     // 评标表id
	ArticleId    int                  `json:"article_id"`             // 所属文章id
	Uid          string               `json:"uid"`                    // 评论人uid
	UidContainer string               `json:"uid_container"`          // 评论人内容
	UidLike      int                  `json:"uid_like"`               // 评论人被点赞数
	Reply        *ArticleCommentReply `orm:"rel(fk)"`                 //设置一对多关系
	Date         time.Time            `orm:"auto_now_add;type(date)"` // 用户创建时间
}

type ArticleCommentReply struct { // 回复评论表
	Id           int       `json:"id"`                     // 评标表id
	ArticleId    int       `json:"article_id"`             // 所属文章id
	ReplyUserUid string    `json:"reply-user-uid"`         // 被回复的评论人uid
	Uid          string    `json:"uid"`                    // 评论人uid
	UidContainer string    `json:"uid_container"`          // 评论人内容
	UidLike      int       `json:"uid_like"`               // 评论人被点赞数
	Date         time.Time `orm:"auto_now_add;type(date)"` // 用户创建时间
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
		// orm.RegisterModel(new(Users), new(MaoShanUser)) // 注册数据表
	}
	//关闭数据库
	defer conn.Close()

}
