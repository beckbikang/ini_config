
功能

	a simple ini file parser
	一个简单的ini文件的解析器


ini文件(ini file)

```

;global-data
[global]

pid = run/php-fpm.pid
;a
error_log = /data1/logs/php-fpm_error.log
;b
log_level = notice
emergency_restart_threshold = 10
emergency_restart_interval = 1m
process_control_timeout = 5s
daemonize = yes
```


简单的使用-simple use

```
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
```

代码输出 output
```
/data1/log/php-fpm_error.log
notice
10
1m
5s
yes
;global-data
;a
/data1/log/php-fpm_error.log
notice
10
1m
5s
yes
```





单元测试 unittest benchmarktest

```
BenchmarkParserConfig-4           	  100000	     12054 ns/op
BenchmarkParserConfig2-4          	  100000	     11961 ns/op
BenchmarkParserConfigParallel-4   	  200000	      7165 ns/op
PASS
ok  	config	4.174s
```