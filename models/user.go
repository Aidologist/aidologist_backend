package models

import "wkBackEnd/utils/databases"

//========================================
// User model
//========================================

type User struct {
	Id       int
	EMail 	 string
	Username string
	Password string
	//Profile  *Profile `orm:"null;rel(one);on_delete(set_null)"`
}

func (this *User) AddUser() (int64, bool) {
	success, o := databases.ConnectOrm()
	if success {
		index, err := o.Insert(this)
		if err == nil {
			return index, true
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
}

func (this *User) GetUser() bool {
	success, o := databases.ConnectOrm()
	if success {
		err := o.Read(this)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
//========================================
// Profile model
//========================================

//type Profile struct {
//	Id      int
//	Gender  string
//	Age     int
//	Address string
//	Email   string
//	User 	*User     `orm:"null;reverse(one)"`
//}