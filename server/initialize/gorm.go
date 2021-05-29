package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"myblog/global"
	"myblog/model"
	"os"
	"strings"
)

func Gorm() *gorm.DB {
	m := global.GCONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Dsn()
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   //连接串
		DefaultStringSize:         191,   //string类型字段的默认长度
		DisableDatetimePrecision:  true,  //禁用datetime精度
		DontSupportRenameIndex:    true,  //重命名索引是采用删除并新建的方式
		DontSupportRenameColumn:   true,  //用change重命名列
		SkipInitializeWithVersion: false, //根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.GLOG.Error("init database failed", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch strings.ToLower(global.GCONFIG.Mysql.LogZap) {
	case "silent":
		config.Logger = Default.LogMode(logger.Silent)
	case "error":
		config.Logger = Default.LogMode(logger.Error)
	case "warn":
		config.Logger = Default.LogMode(logger.Warn)
	case "info":
		config.Logger = Default.LogMode(logger.Info)
	case "zap":
		config.Logger = Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = Default.LogMode(logger.Info)
			break
		}
		config.Logger = Default.LogMode(logger.Silent)
	}
	return config
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Article{},
		model.Category{},
		model.Comments{},
	)
	if err != nil {
		global.GLOG.Error("register tables failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GLOG.Info("register tables successfully")
}
