package tcpbridge

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var logger *log.Logger
var once sync.Once

func init() {
	NewLogger()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetLevel(log.ErrorLevel)

}

func NewLogger() *log.Logger {
	once.Do(func() {
		infoPath := "logs/info.log"
		writerInfo, _ := rotatelogs.New(
			infoPath+".%Y%m%d",
			rotatelogs.WithLinkName(infoPath),
			rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
			rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
		)

		errorPath := "logs/error.log"
		writerError, _ := rotatelogs.New(
			errorPath+".%Y%m%d",
			rotatelogs.WithLinkName(errorPath),
			rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
			rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
		)
		logger = log.New()

		logger.Hooks.Add(lfshook.NewHook(
			lfshook.WriterMap{
				log.InfoLevel:  writerInfo,
				log.ErrorLevel: writerError,
			},
			&log.JSONFormatter{},
		))
	})
	return logger
}
