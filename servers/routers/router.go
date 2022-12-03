package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"servers/controllers"
	"servers/controllers/maoshanuser"
	"servers/controllers/messageforme"
	"servers/controllers/user"
	"servers/controllers/xuankonguser"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/statistics", &controllers.StatisticsController{})
	beego.Router("/user/sign_in", &user.UsersController{}, "get:Get")                                     // 登录
	beego.Router("/user/sign_up", &user.UsersController{}, "post:Post")                                   // 注册
	beego.Router("/user/save_maoshanuser", &maoshanuser.MaoshanuserController{}, "post:Post;get:Get")     // 奇门保存查看卦
	beego.Router("/user/save_xuankonguser", &xuankonguser.XuanKongUsersController{}, "post:Post;get:Get") // 玄空保存查看卦
	beego.Router("messageforme", &messageforme.MessageForMeController{}, "post:Post;get:Get")             // 玄空保存查看卦
}
