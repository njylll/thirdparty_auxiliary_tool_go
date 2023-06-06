package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime"
)

const (
	logPath = "./log/thirdparty_auxiliary_tool.log"
)

// InitLog 初始化日志
func Setup() {
	logrus.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    30, // megabytes
		MaxBackups: 3,
		MaxAge:     7, //days
	}))
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05",
		CustomCallerFormatter: func(f *runtime.Frame) string {
			return fmt.Sprintf(" [%s:%d]\t", path.Base(f.File), f.Line)
		},
	})

	return
}
