package logger

import (
	"fmt"
	"github.com/iiran/lltt/pkg/setting"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F          *os.File
	logger     *log.Logger
	logPrefix  string
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	LinePrefix string
	TraceDepth int
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup(cfg *setting.ServerConfigLog) {
	LogSavePath = cfg.PathName
	LogSaveName = cfg.FileName
	LogFileExt = cfg.FileExt
	LinePrefix = cfg.LinePrefix
	TraceDepth = cfg.TraceDepth

	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, LinePrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(TraceDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s] Get Caller failed", levelFlags[level])
	}

	logger.SetPrefix(LinePrefix + logPrefix)
}
