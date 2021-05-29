package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"myblog/config"
	"time"
)

var (
	GDB     *gorm.DB
	GREDIS  *redis.Client
	GCONFIG config.Server
	GVP     *viper.Viper
	GLOG    *zap.Logger
)

type GMODEL struct {
	ID        uint `gorm:"primaryKey;unique;autoIncrement"`
	CreateAt  time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt
}
