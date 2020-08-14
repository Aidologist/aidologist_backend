package modelsFunc

import "github.com/astaxie/beego/orm"

func ConnectORM() (bool, orm.Ormer)  {
	o := orm.NewOrm()
	if o.Using("xq") == nil {
		return true, o
	} else {
		return false, o
	}
}