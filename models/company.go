package models

import "wkBackEnd/utils/databases"

type Company struct {
	Id       int
	EMail 	 string
	Name	 string
	Password string
}

func (this *Company) AddCompany() (int64, bool) {
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

func (this *Company) GetCompany() bool {
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