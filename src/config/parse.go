package config

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

//将数据写入到文件, 无法保证配置的顺序
func SaveConfigToFile(cf *IniConfig, filename string) bool {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("open failed")
		return false
	}
	defer fp.Close()

	for secName, sections := range cf.Sections {
		comment := cf.GetConfigComment(secName)
		if len(comment) != 0 {
			fp.WriteString(comment + "\n")
		}
		fp.WriteString("[" + secName + "]\n")
		for _, val := range sections {
			commentData := cf.GetConfigCommentData(secName, val.key)
			if len(commentData) != 0 {
				fp.WriteString(commentData + "\n")
			}
			fp.WriteString(val.key + "=" + val.val + "\n")
		}
	}

	return true
}

//解析config
func ParserConfig(filename string, reload bool) (*IniConfig, error) {
	if !reload {
		if cf, ok := configList[filename]; ok {
			return cf, nil
		}
	}
	return parserFile(filename)
}

//真正的解析config的文件
func parserFile(filename string) (*IniConfig, error) {

	acf, err := NewIniConfig()
	if err != nil {
		return nil, err
	}

	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	bufRead := bufio.NewReader(fp)

	//记录注释
	isComment := false

	var curSectionName string
	var comment string
	curSectionName = ""
	for {

		line, err := bufRead.ReadBytes('\n')
		if len(line) == 0 {
			break
		}
		//去除空格
		line = bytes.TrimSpace(line)
		//去除空行
		if len(line) == 0 {
			continue
		}
		//记录注释，注释的下一个就是需要存放的地方
		if line[0] == ';' {
			if !isComment {
				isComment = true
				comment = strings.TrimSpace(string(line)) + "\n"
			} else {
				comment = comment + string(line)
			}
		} else if line[0] == '[' {
			line = bytes.Trim(line, "[")
			line = bytes.Trim(line, "]")
			curSectionName = string(line)
			//存在注释，记录注释
			if isComment {
				commentName := getSectionCommentName(curSectionName)
				if _, ok := acf.Comment[curSectionName]; !ok {
					acf.Comment[curSectionName] = make(map[string]*commentConfig)
				}
				acf.Comment[curSectionName][commentName] = NewCommentConfig(commentName, comment)
				isComment = false
			}
		} else {
			if len(curSectionName) == 0 {
				return nil, FORMAT_ERROR
			}
			//处理内容
			lineString := string(line)
			lineArr := strings.Split(lineString, "=")

			if len(lineArr) == 2 {
				dataKey := strings.TrimSpace(lineArr[0])
				dataVal := strings.TrimSpace(lineArr[1])
				//获取内容
				if _, ok := acf.Sections[curSectionName]; !ok {
					acf.Sections[curSectionName] = make(map[string]*sectionConfig)
				}
				acf.Sections[curSectionName][dataKey] = NewSection(dataKey, dataVal)
				if isComment {
					acf.Comment[curSectionName][dataKey] = NewCommentConfig(dataKey, comment)
				}
				//log.Println(acf.Sections[curSectionName][dataKey].key)
			}

			isComment = false
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

	}
	return acf, nil
}
