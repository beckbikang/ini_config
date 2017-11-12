package config

import (
	"testing"
)

//测试获取数据
func TestParserConfig(t *testing.T) {
	//测试获取数据
	cf, err := ParserConfig("../config_file/simple.conf", false)

	if err != nil {
		t.Fatal("get config faild", err)
		return
	}

	//获取数据
	if cf.GetConfig("global", "log_level") != "notice" {
		t.Error("get log level faild")
	}
	if cf.GetConfigInt("global", "emergency_restart_threshold") != 10 {
		t.Error("get log level faild")
	}

	//获取评论
	if cf.GetConfigCommentData("global", "log_level") != ";b" {
		t.Error("get global log_level comment faild")
	}
}

func TestSaveConfigToFile(t *testing.T) {

	cf, err := ParserConfig("../config_file/simple.conf", false)
	if err != nil {
		t.Fatal("get config faild", err)
		return
	}

	SaveConfigToFile(cf, "../config_file/simple_test.conf")

	cf2, err := ParserConfig("../config_file/simple_test.conf", false)
	if err != nil {
		t.Fatal("get config faild", err)
		return
	}
	//获取数据
	if cf2.GetConfig("global", "log_level") != "notice" {
		t.Error("get log level faild")
	}
	if cf2.GetConfigInt("global", "emergency_restart_threshold") != 10 {
		t.Error("get log level faild")
	}
	//获取评论
	if cf2.GetConfigCommentData("global", "log_level") != ";b" {
		t.Error("get global log_level comment faild")
	}
	return
}

//解析性能测试
func BenchmarkParserConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParserConfig("../config_file/simple.conf", false)
	}
}

func BenchmarkParserConfig2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParserConfig("../config_file/simple.conf", true)
	}
}

// 测试并发效率
func BenchmarkParserConfigParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ParserConfig("../config_file/simple.conf", false)
		}
	})
}
