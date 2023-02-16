package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/controllers"
	"servers/controllers/article"
	"servers/controllers/articlecomment"
	"servers/controllers/articlecommentreply"
	"servers/controllers/articlerecommendation"
	"servers/controllers/maoshanuser"
	"servers/controllers/messageforme"
	"servers/controllers/updatacover"
	"servers/controllers/user"
	"servers/controllers/userinfo"
	"servers/controllers/xuankonguser"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/statistics", &controllers.StatisticsController{})
	beego.Router("/user/sign_in", &user.UsersController{}, "get:Get")                                       // 登录
	beego.Router("/user/sign_up", &user.UsersController{}, "post:Post")                                     // 注册
	beego.Router("/user/save_maoshanuser", &maoshanuser.MaoshanuserController{}, "post:Post;get:Get")       // 奇门保存查看卦
	beego.Router("/user/save_xuankonguser", &xuankonguser.XuanKongUsersController{}, "post:Post;get:Get")   // 玄空保存查看卦
	beego.Router("/messageforme", &messageforme.MessageForMeController{}, "post:Post;get:Get")              // 给我留言接口
	beego.Router("/article", &article.ArticleController{}, "post:Post;get:Get")                             // 文章发布接口
	beego.Router("/article_recommendation", &articlerecommendation.ArticleRecommendation{}, "get:Get")      // 文章推荐查询接口
	beego.Router("/article_comment", &articlecomment.ArticleComment{}, "get:Get;post:Post")                 // 一级评论发布，查询接口
	beego.Router("/article_comment_reply", &articlecommentreply.ArticleCommentReply{}, "get:Get;post:Post") // 二级评论发布，查询接口
	beego.Router("/update_cover", &updatacover.UpdataCover{}, "get:Get;post:Post;delete:Delete")            // 图片上传接口
	beego.Router("/user/info", &userinfo.InfoController{}, "get:Get;post:Post;delete:Delete")               // 用户信息查询
}
