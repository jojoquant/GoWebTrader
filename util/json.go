package util

import (
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
)

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (j *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("json 文件读取错误, err:", err)
		return
	}

	err = jsoniter.Unmarshal(data, v)
	if err != nil {
		log.Println("json 文件解码错误, err:", err)
		return
	}
}
