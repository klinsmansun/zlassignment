package std

import (
	"log"
	"os"

	"github.com/klinsmansun/zlassignment/model"
)

type logger struct {
	logLevel int
	instance *log.Logger
}

func CreateLogger(logLevel int) model.LogUsecase {
	l := &logger{
		logLevel: logLevel,
		instance: log.New(os.Stdout, "", 0),
	}

	return l
}

func (l *logger) LogEmerg(msg ...interface{}) {
	l.LogMessage(model.LoglevelEmerg, msg...)
}

func (l *logger) LogAlert(msg ...interface{}) {
	l.LogMessage(model.LoglevelAlert, msg...)
}

func (l *logger) LogCrit(msg ...interface{}) {
	l.LogMessage(model.LoglevelCrit, msg...)
}

func (l *logger) LogErr(msg ...interface{}) {
	l.LogMessage(model.LoglevelErr, msg...)
}

func (l *logger) LogWarning(msg ...interface{}) {
	l.LogMessage(model.LoglevelWarning, msg...)
}

func (l *logger) LogNotice(msg ...interface{}) {
	l.LogMessage(model.LoglevelNotice, msg...)
}

func (l *logger) LogInfo(msg ...interface{}) {
	l.LogMessage(model.LoglevelInfo, msg...)
}

func (l *logger) LogDebug(msg ...interface{}) {
	l.LogMessage(model.LoglevelDebug, msg...)
}

func (l *logger) LogMessage(targetLevel int, msg ...interface{}) {
	if l.logLevel >= targetLevel {
		l.instance.Println(msg...)
	}
}
