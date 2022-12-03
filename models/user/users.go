package user

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"servers/models"
)

// QueryUserModelData uid查询账号信息
func QueryUserModelData(uid string) (data []*models.Users, error error) {
	o := orm.NewOrm()
	var user []*models.Users
	err := o.QueryTable("users").Filter("uid", uid).One(&user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// SignInUserModelData id查询账号信息
func SignInUserModelData(account string, password string) (data []*models.Users, error error) {
	o := orm.NewOrm()
	var user []*models.Users
	err := o.QueryTable("users").Filter("account", account).One(&user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// InsertUserModelData 插入账号信息
func InsertUserModelData(account string, password string, name string, uid string) (id string, error error) {
	o := orm.NewOrm()
	var insertUsers models.Users
	insertUsers.Account = account
	insertUsers.Password = password
	insertUsers.Name = name
	insertUsers.Uid = uid

	_, error = o.Insert(&insertUsers)
	if error != nil {
		fmt.Println(error)
		return "", error
	}
	return uid, error
}

// DeduplicationUser 注册账号去重
func DeduplicationUser(account string) (flag bool, error error) {
	o := orm.NewOrm()

	var user []*models.Users
	err := o.QueryTable("users").Filter("account", account).One(&user)
	if err != nil {
		return false, nil
	}
	return true, err
	//
	//var deduplicationUsers models.Users
	//deduplicationUsers.Account = account
	//
	//err := o.QueryTable("users").Filter("account", account).One(&deduplicationUsers)
	//if err == orm.ErrMultiRows {
	//	// 多条的时候报错
	//	return true, err
	//}
	//if err == orm.ErrNoRows {
	//	// 没有找到记录
	//	return false, err
	//}
	//if err != nil {
	//	return false, nil
	//}
	//return true, nil
}
