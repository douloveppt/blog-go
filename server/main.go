package main

import (
	"myblog/core"
	"myblog/global"
	"myblog/initialize"
)

func main() {
	global.GVP = core.Viper()
	global.GLOG = core.Zap()
	global.GDB = initialize.Gorm()
	if global.GDB != nil {
		initialize.MysqlTables(global.GDB)
		db, _ := global.GDB.DB()
		defer db.Close()
	} else {
		global.GLOG.Panic("init database failed")
	}
}
