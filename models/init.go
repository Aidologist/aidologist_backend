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
	//========== Chat room service ==========
	orm.RegisterModel(new(MessageTrackStatus), new(MessageTrack), new(Message))
	orm.RegisterModel(new(ChatRoom))
	orm.RegisterModel(new(MessageInTrack))
	orm.RegisterModel(new(UserInChatRoom))
	//========== Main service ==========
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(Company))
	orm.RegisterModel(new(CompanyFavoriteUser))   // m2m tables needed to be added here
	orm.RegisterModel(new(UserFavoriteTask))   // m2m tables needed to be added here
	//========================================
	// Run model registers
	//========================================
	_ = orm.RunSyncdb("default", false, true)
}