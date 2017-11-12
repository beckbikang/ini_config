


简单的使用

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
}









单元测试


BenchmarkParserConfig-4           	  100000	     12054 ns/op
BenchmarkParserConfig2-4          	  100000	     11961 ns/op
BenchmarkParserConfigParallel-4   	  200000	      7165 ns/op
PASS
ok  	config	4.174s