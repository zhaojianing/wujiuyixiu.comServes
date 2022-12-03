package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	// 自动建表
	//orm.RunSyncdb("default", false, true)
	//注册 model
	fmt.Println("更新啦！！！！！！！！！！更新必须把更新的表写在最前面，否则不更新")
	orm.RegisterModel(new(XuanKongUser), new(MessageForMe), new(Users), new(MaoShanUser), new(Statistics)) // 注册数据表
}

func UpdateStatistics() (nun int, e error) {
	o := orm.NewOrm()
	statistics := Statistics{Id: 1}
	if o.Read(&statistics) == nil {
		statistics.Totalpeople = statistics.Totalpeople + 1
		if num, err := o.Update(&statistics); err == nil {
			fmt.Println(num)
			return statistics.Totalpeople, nil
		}
	}
	return 0, e
}
