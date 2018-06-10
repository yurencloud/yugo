package config

import (
	"strings"
	"os"
	"bufio"
	"io"
)

// 读取所有config文件，default.conf > app.conf > other.conf 并合并字符串
type Config struct {
	Map map[string]string
}

var DefaultConfig  = map[string]string {
	"port" : "3000",
	"static" : "static",
	"active" : "app",
	"template.path" : "./static/views",
	"template.suffix" : ".jet",
	"csrf.enabled" : "true",
	"csrf.key" : "5ebe2294ecd0e0f08eab7690d2a6ee69",
	"csrf.request.header" : "csrf",
	"csrf.field.name" : "csrf",
	"csrf.max.age" : "43200",
	"session.key" : "5ebe2294ecd0e0f08eab7690d2a6ee69",
	"app.name" : "yugo",
	"log.max.size" : "10485760",
	"log.level" : "debug",
}

func ReadAllConfigFile() Config {
	// 先读取默认配置
	config := Config{DefaultConfig }

	// 循环读取app,dev,prod等conf文件,并覆盖默认配置
	for {
		activeConfig := ReadConfigFile("./config/" + config.Map["active"] + ".conf")
		for key, value := range activeConfig.Map {
			config.Map[key] = value
		}
		if len(activeConfig.Map["active"]) == 0 {
			break
		}
	}

	return config
}

func ReadConfigFile(path string) Config {
	var config Config

	config.Map = make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// 无限循环，直到break
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(line))

		// 跳过开头是#的注释
		if strings.Index(s, "#") == 0 {
			continue
		}

		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		first := strings.TrimSpace(s[:index])
		if len(first) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		// 处理各种注释
		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		config.Map[first] = strings.TrimSpace(second)
	}

	return config
}

func Get(key string) string {
	config := ReadAllConfigFile()
	return config.Map[key]
}

func GetConfigMap() map[string]string {
	config := ReadAllConfigFile()
	return config.Map
}
