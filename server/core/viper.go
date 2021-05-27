package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	conf "myblog/config"
	"myblog/global"
	"os"
)

// Viper 初始化配置文件，优先级：Viper参数 > 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(conf.Env); configEnv == "" {
				config = conf.File
				fmt.Printf("chosen default config file: %v\n", conf.File)
			} else {
				config = configEnv
				fmt.Printf("chosen config from environment: %v\n", config)
			}
		} else {
			fmt.Printf("chosen config from command line: %v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("chosen config from func Viper'param: %v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err := v.Unmarshal(&global.GCONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GCONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
