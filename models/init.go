package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"wkBackEnd/utils/constants"
	"wkBackEnd/utils/databases"
)

func init()  {
	//========================================
	// MYSQL DATABASE model registers
	//========================================
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// registers of default database
	_ = orm.RegisterDataBase("default", "mysql", databases.MysqlConnectionString(
		constants.MysqlUsername,
		constants.MysqlPassword,
		constants.MysqlRegisterDataBaseName,
		constants.MysqlCharset,
		constants.MysqlUrl,
		constants.MysqlPort))
	// registers of other databases
	_ = orm.RegisterDataBase(constants.MysqlRegisterDataBaseName, "mysql", databases.MysqlConnectionString(
		constants.MysqlUsername,
		constants.MysqlPassword,
		constants.MysqlRegisterDataBaseName,
		constants.MysqlCharset,
		constants.MysqlUrl,
		constants.MysqlPort))
	//========================================
	// Registe data models
	//========================================
	//========== Main service ==========

	//========================================
	// Run model registers
	//========================================
	_ = orm.RunSyncdb("default", false, true)
}


















