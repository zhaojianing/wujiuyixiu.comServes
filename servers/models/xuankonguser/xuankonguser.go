package xuankonguser

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func InsertXuanKonguser(uid string, time string, title string, introduction string, isPublic int) (id int64, error error) {
	o := orm.NewOrm()
	var xuankonguser models.XuanKongUser
	xuankonguser.Uid = uid
	xuankonguser.Time = time
	xuankonguser.Title = title
	xuankonguser.Introduction = introduction
	xuankonguser.IsPublic = isPublic

	id, error = o.Insert(&xuankonguser)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	return id, error
}

func GetXuanKonguser(uid string) (data []*models.XuanKongUser, error error) {
	o := orm.NewOrm()
	var dataMaoShan []*models.XuanKongUser
	_, err := o.QueryTable("XuanKongUser").Filter("uid", uid).All(&dataMaoShan)
	if err != nil {
		return nil, err
	}
	return dataMaoShan, nil
}
