package model

const (
	LoglevelEmerg int = iota
	LoglevelAlert
	LoglevelCrit
	LoglevelErr
	LoglevelWarning
	LoglevelNotice
	LoglevelInfo
	LoglevelDebug
)

type LogUsecase interface {
	LogEmerg(msg ...interface{})
	LogAlert(msg ...interface{})
	LogCrit(msg ...interface{})
	LogErr(msg ...interface{})
	LogWarning(msg ...interface{})
	LogNotice(msg ...interface{})
	LogInfo(msg ...interface{})
	LogDebug(msg ...interface{})
}
