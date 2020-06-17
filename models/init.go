package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"wkBackEnd/utils/constants"
	"wkBackEnd/utils/strings"
)

func init()  {
	//========================================
	// MYSQL DATABASE model registers
	//========================================
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// registers of default database
	_ = orm.RegisterDataBase("default", "mysql", strings.MysqlConnectionString(
		constants.MysqlUsername,
		constants.MysqlPassword,
		constants.MysqlRegisterDataBaseName,
		constants.MysqlCharset,
		constants.MysqlUrl,
		constants.MysqlPort))
	// registers of other databases
	_ = orm.RegisterDataBase(constants.MysqlRegisterDataBaseName, "mysql", strings.MysqlConnectionString(
		constants.MysqlUsername,
		constants.MysqlPassword,
		constants.MysqlRegisterDataBaseName,
		constants.MysqlCharset,
		constants.MysqlUrl,
		constants.MysqlPort))
	//========================================
	// Registe data models
	//========================================
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Task))
	//========================================
	// Run model registers
	//========================================
	_ = orm.RunSyncdb("default", false, true)
}