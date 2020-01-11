package logger

import (
	"fmt"
	"github.com/iiran/lltt/pkg/core"
	"github.com/iiran/lltt/pkg/helper"
	"log"
	"os"
	"time"
)

var (
	LogSavePath string
	LogSaveName string
	LogFileExt  string

)

func getLogPath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogPath()
	suffixPath := fmt.Sprintf("%s-%s.%s", LogSaveName, time.Now().Format(core.TIME_FORMAT), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(path2file string) *os.File {
	_, err := os.Stat(path2file)
	switch {
	case os.IsNotExist(err):
		helper.MakeDir(getLogPath())
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(path2file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}
