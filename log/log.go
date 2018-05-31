package log


import (
	"os"
	log "github.com/Sirupsen/logrus"
	"yugo/config"
	"strconv"
	"time"
	"fmt"
	"io"
)

func init() {
	appName := config.Get("app.name")
	level  := config.Get("log.level")
	logMaxSize,_ := strconv.ParseInt(config.Get("log.max.size"), 10, 64)
	logName := "./log/"+appName+".log";

	logFile, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)

	if err!=nil {
		fmt.Println(err)
	}

	file, err := os.Stat(logName)

	// 如果大于10M,就重命名，并创建
	if  file.Size() > logMaxSize {
		timeString := strconv.FormatInt(time.Now().Unix(),10)
		os.Rename(logName, "./log/"+appName+"-"+timeString+".log")
		logFile, err = os.OpenFile(logName, os.O_CREATE, 0755)

		if err!=nil {
			fmt.Println(err)
		}
	}

	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)

	logLevel := map[string]log.Level{
		"panic":log.PanicLevel,
		"fatal":log.FatalLevel,
		"error":log.ErrorLevel,
		"warning":log.WarnLevel,
		"info":log.InfoLevel,
		"debug":log.DebugLevel,
	}

	// Only log the warning severity or above.
	log.SetLevel(logLevel[level])
}