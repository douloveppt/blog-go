package main

import (
	"fmt"
	"myblog/core"
	"myblog/global"
)

func main() {
	global.GVP = core.Viper()
	fmt.Println(global.GVP)
}
