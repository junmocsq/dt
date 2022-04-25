package derror

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func NewErr(code int) error {
	return errors.New(Message(code))
}

func Message(code int) string {
	cpdeMap, err := CodeMapping()
	if err != nil {
		return "未知错误"
	}
	str := cpdeMap[code]
	if str == "" {
		return "不存在的code码"
	}
	return cpdeMap[code]
}

func CodeMapping() (map[int]string, error) {
	codeMap := make(map[int]string)
	os.Open("err.json")
	file, err := os.Open("derror/err.json")
	if err != nil {
		return codeMap, err
	}
	defer func() {
		file.Close()
	}()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return codeMap, err
	}
	json.Unmarshal(b, &codeMap)
	return codeMap, err
}
