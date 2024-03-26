package log

import (
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

var logger *log.Logger

func NewLogger(isDebug bool, filePath string) {
	if logger == nil {
		logger = log.New()
	}
	if isDebug {
		logger.Level = log.DebugLevel
	} else {
		logger.Level = log.InfoLevel
	}

	// 命令行输出格式
	loggerFormatter := &prefixed.TextFormatter{
		DisableTimestamp: true,
		ForceFormatting:  true,
		ForceColors:      true,
		DisableColors:    false,
	}

	logger.SetFormatter(loggerFormatter)
	logger.SetOutput(os.Stdout)

	// 日志输出到文件
	if filePath != "" {
		fileFormatter := &prefixed.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02.15:04:05.000000",
			ForceFormatting: true,
			ForceColors:     false,
			DisableColors:   true,
		}

		pathMap := lfshook.PathMap{
			log.InfoLevel:  filePath,
			log.DebugLevel: filePath,
			log.ErrorLevel: filePath,
		}

		newHooks := make(log.LevelHooks)
		newHooks.Add(lfshook.NewHook(
			pathMap,
			fileFormatter,
		))

		//logger.AddHook(lfshook.NewHook()) // 使用 Replace 而不使用 Add
		logger.ReplaceHooks(newHooks)
	}
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}
