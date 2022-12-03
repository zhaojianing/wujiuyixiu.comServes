package maoshanuser

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func InsertMaoshanuser(uid string, time string, title string, introduction string, isPublic int) (id int64, error error) {
	o := orm.NewOrm()
	var maoshanuser models.MaoShanUser
	maoshanuser.Uid = uid
	maoshanuser.Time = time
	maoshanuser.Title = title
	maoshanuser.Introduction = introduction
	maoshanuser.IsPublic = isPublic

	id, error = o.Insert(&maoshanuser)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	return id, error
}

func GetMaoshanuser(uid string) (data []*models.MaoShanUser, error error) {
	o := orm.NewOrm()
	var dataMaoShan []*models.MaoShanUser
	_, err := o.QueryTable("MaoShanUser").Filter("uid", uid).All(&dataMaoShan)
	if err != nil {
		return nil, err
	}
	return dataMaoShan, nil
}
