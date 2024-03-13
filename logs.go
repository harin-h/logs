package logs

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewDevelopmentConfig()
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Error(msg interface{}, fields ...zap.Field) {
	switch v := msg.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
