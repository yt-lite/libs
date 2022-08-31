package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	var err error
	config := zap.NewProductionConfig()

	enccoderConfig := zap.NewProductionEncoderConfig()
	enccoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	enccoderConfig.StacktraceKey = "" // to hide stacktrace info
	config.EncoderConfig = enccoderConfig

	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	sugar = zapLog.Sugar()
}

func Info(message string, fields interface{}) {
	sugar.Info(message, fields)
}

func Infof(message string, fields interface{}) {
	sugar.Infof(message, fields)
}

func Infoln(message string, fields interface{}) {
	sugar.Infoln(message, fields)
}

func Warnln(message string, fields interface{}) {
	sugar.Warnln(message, fields)
}

func Warnf(message string, args ...interface{}) {
	sugar.Warnf(message, args...)
}

func Warn(message string, fields interface{}) {
	sugar.Warn(message, fields)
}

func Debug(message string, fields interface{}) {
	sugar.Debug(message, fields)
}

func Debugf(message string, fields interface{}) {
	sugar.Debugf(message, fields)
}

func Errorf(message string, fields interface{}) {
	sugar.Errorf(message, fields)
}

func Errorln(message string, fields interface{}) {
	sugar.Errorln(message, fields)
}

func Error(message string, fields interface{}) {
	sugar.Error(message)
}

func Fatal(message string, fields interface{}) {
	sugar.Fatal(message, fields)
}

func Fatalf(message string, fields interface{}) {
	sugar.Fatalf(message, fields)
}
func Fatalln(message string, fields interface{}) {
	sugar.Fatalln(message, fields)
}

func Panicf(message string, fields interface{}) {
	sugar.Panicf(message, fields)
}

func Panicln(message string, fields interface{}) {
	sugar.Panicln(message, fields)
}

func Panic(message string, fields interface{}) {
	sugar.Panic(message, fields)
}
