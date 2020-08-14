package databases

import (
	"github.com/astaxie/beego/orm"
	"wkBackEnd/utils/constants"
)

func ConnectOrm() (bool, orm.Ormer) {
	o := orm.NewOrm()
	if o.Using(constants.MysqlRegisterDataBaseName) == nil {
		return true, o
	} else {
		return false, o
	}
}