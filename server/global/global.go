package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"myblog/config"
)

var (
	GDB     *gorm.DB
	GREDIS  *redis.Client
	GCONFIG config.Server
	GVP     *viper.Viper
	GLOG    *zap.Logger
)
