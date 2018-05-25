package config

import (
	"io/ioutil"
	"fmt"
)

func ReadConfig() (string, error) {
	if configText, err := ioutil.ReadFile("./confi/default.conf"); err != nil {
		fmt.Printf("文件不存在: %s\n", err)
		return "", fmt.Errorf("error %q",err)
	}else{
		return string(configText), nil
	}
}
