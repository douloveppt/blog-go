package main

import (
	"myblog/core"
	"myblog/global"
	"time"
)

func main() {
	global.GVP = core.Viper()
	global.GLOG = core.Zap()
	for {
		global.GLOG.Debug("this is a debug log")
		global.GLOG.Info("this is a ingo log")
		global.GLOG.Warn("this is a warn log")
		global.GLOG.Error("this is a error log")
		global.GLOG.DPanic("this is a dpanic log")
		//global.GLOG.Panic("this is a panic log")
		//global.GLOG.Fatal("this is a fatal log")
		time.Sleep(time.Second)
	}
}
