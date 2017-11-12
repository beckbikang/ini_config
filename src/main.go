package main

import (
	"config"
	"fmt"
)

func main() {
	t_iniconfig()
}

func t_iniconfig() {
	cf, err := config.ParserConfig("./config_file/simple.conf", false)
	if err != nil {
		fmt.Println("get config faild")
		return
	}
	//获取配置
	fmt.Println(cf.GetConfig("global", "error_log"))
	fmt.Println(cf.GetConfig("global", "log_level"))
	fmt.Println(cf.GetConfig("global", "emergency_restart_threshold"))
	fmt.Println(cf.GetConfig("global", "emergency_restart_interval"))
	fmt.Println(cf.GetConfig("global", "process_control_timeout"))
	fmt.Println(cf.GetConfig("global", "daemonize"))

	//获取注释
	fmt.Println(cf.GetConfigComment("global"))
	fmt.Println(cf.GetConfigCommentData("global", "error_log"))

	//将配置写入到新的文件
	config.SaveConfigToFile(cf, "./config_file/simple_bak.conf")

	cf, err = config.ParserConfig("./config_file/simple_bak.conf", false)
	if err != nil {
		fmt.Println("get config faild")
		return
	}

	fmt.Println(cf.GetConfig("global", "error_log"))
	fmt.Println(cf.GetConfig("global", "log_level"))
	fmt.Println(cf.GetConfig("global", "emergency_restart_threshold"))
	fmt.Println(cf.GetConfig("global", "emergency_restart_interval"))
	fmt.Println(cf.GetConfig("global", "process_control_timeout"))
	fmt.Println(cf.GetConfig("global", "daemonize"))

}
