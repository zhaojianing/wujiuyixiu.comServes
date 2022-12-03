package messageforme

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

func InsertMessageForMe(uid string, name string, introduction string, date string) (id int64, error error) {
	o := orm.NewOrm()
	var messageforme models.MessageForMe
	messageforme.Uid = uid
	messageforme.Name = name
	messageforme.Introduction = introduction
	messageforme.Date = date

	id, error = o.Insert(&messageforme)
	if error != nil {
		fmt.Println(error)
		return 0, error
	}
	return id, error
}

func GetMessageForMe(pagination int) (data []*models.MessageForMe, error error) {
	o := orm.NewOrm()
	var dataMessageForMes []*models.MessageForMe
	_, err := o.QueryTable("MessageForMe").OrderBy("-date", "-id").Limit(10, (pagination-1)*10).All(&dataMessageForMes)
	if err != nil {
		return nil, err
	}
	return dataMessageForMes, nil
}

func GetMessageForMeCount() (count int64, error error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("MessageForMe").Count() // SELECT COUNT(*) FROM USER
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
