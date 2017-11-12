package config

import (
	"errors"
	"fmt"
	//"log"
	"strconv"
	"strings"
)

const (
	DEFAULT_NAME   = "default_key"
	SectionComment = "comment-section-%s"
)

var (
	FORMAT_ERROR = errors.New("ini format error")
	NOT_FOUND    = errors.New("not found config")
)

type IniConfig struct {
	Sections map[string]sectionConfigs
	Comment  map[string]commentConfigs //注释
}

//注释
type commentConfig struct {
	name    string
	comment string
}
type commentConfigs map[string]*commentConfig
type sectionConfig struct {
	key string
	val string //需要的时候可以去转换
}
type sectionConfigs map[string]*sectionConfig

//创建一个configlist
var configList = make(map[string]*IniConfig)

//创建一个空的config
func NewIniConfig() (*IniConfig, error) {
	acf := &IniConfig{}
	acf.Sections = make(map[string]sectionConfigs)
	acf.Comment = make(map[string]commentConfigs)

	return acf, nil
}

//获取配置
func (cf *IniConfig) GetConfig(secting, name string) string {
	sec, ok := cf.Sections[secting][name]
	if !ok {
		return ""
	}
	return sec.val
}

//获取配置的各种方式 int
func (cf *IniConfig) GetConfigInt(secting, name string) int {
	ret := cf.GetConfig(secting, name)
	if len(ret) == 0 {
		return 0
	}
	i, err := strconv.Atoi(ret)
	if err != nil {
		return 0
	}
	return i
}

//更新或者创建配置
func (cf *IniConfig) PutConfig(secting, name, val string) bool {
	sec, ok := cf.Sections[secting]
	if !ok {
		cf.Sections[secting] = make(map[string]*sectionConfig)
		cf.Sections[secting][name] = NewSection(name, val)
		return true
	}
	sec[name].key = name
	sec[name].val = val
	return true
}

//删除配置
func (cf *IniConfig) DelConfigData(secting, name string) bool {
	_, ok := cf.Sections[secting][name]
	if !ok {
		return true
	}
	delete(cf.Sections[secting], name)
	return true
}

//获取配置的各种方式 int
func (cf *IniConfig) GetConfigDouble(secting, name string) float64 {
	ret := cf.GetConfig(secting, name)
	if len(ret) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(ret, 64)
	if err != nil {
		return 0
	}
	return f
}

//获取注释数据
func (cf *IniConfig) GetConfigComment(secting string) string {
	name := getSectionCommentName(secting)
	ret, ok := cf.Comment[secting][name]
	if !ok {
		return ""
	}
	return strings.TrimSpace(ret.comment)
}

//获取注释数据
func (cf *IniConfig) GetConfigCommentData(secting, name string) string {
	ret, ok := cf.Comment[secting][name]
	if !ok {
		return ""
	}
	return strings.TrimSpace(ret.comment)
}

//更新注释配置
func (cf *IniConfig) PutConfigCommentData(secting, name, val string) bool {
	_, ok := cf.Comment[secting][name]
	if !ok {
		cf.Comment[secting] = make(map[string]*commentConfig)
		cf.Comment[secting][name] = NewCommentConfig(name, val)
		return true
	}
	cf.Comment[secting][name].name = name
	cf.Comment[secting][name].comment = val
	return true
}

func getSectionCommentName(sectionName string) string {
	return fmt.Sprintf(SectionComment, sectionName)
}

//创建一个comment
func NewCommentConfig(name, comment string) *commentConfig {
	return &commentConfig{name, comment}
}

//创建一个配置
func NewSection(k, v string) *sectionConfig {
	return &sectionConfig{k, v}
}
