package search

import (
	"encoding/json"
	"os"
)

//const 声明常量
const dataFile = "data/data.json"

//Feed 首字母大写 所以该类型对外暴露
//Feed 结构类型
type Feed struct {
	//`` 引号中的部分叫tag(标记)
	Name string `json:"site"`
	ORI  string `json:"link"`
	Type string `json:"type"`
}

//读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}
}
