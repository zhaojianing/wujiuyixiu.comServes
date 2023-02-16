package userinfo

import (
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func GetUser(uid string) (data models.Users, err error) {
	o := orm.NewOrm()
	var user models.Users
	user.Uid = uid

	_, err = o.QueryTable(&user).Filter("Uid", uid).All(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
