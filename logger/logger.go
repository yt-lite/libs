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

func Info(message string) {
	sugar.Info(message)
}

func Infof(message string, fields interface{}) {
	sugar.Infof(message, fields)
}

func Infoln(message string) {
	sugar.Infoln(message)
}

func Warnln(message string) {
	sugar.Warnln(message)
}

func Warnf(message string, args ...interface{}) {
	sugar.Warnf(message, args...)
}

func Warn(message string) {
	sugar.Warn(message)
}

func Debug(message string) {
	sugar.Debug(message)
}

func Debugf(message string, fields interface{}) {
	sugar.Debugf(message, fields)
}

func Errorf(message string, fields interface{}) {
	sugar.Errorf(message, fields)
}

func Errorln(message string) {
	sugar.Errorln(message)
}

func Error(message string) {
	sugar.Error(message)
}

func Fatal(message string) {
	sugar.Fatal(message)
}

func Fatalf(message string, fields interface{}) {
	sugar.Fatalf(message, fields)
}
func Fatalln(message string) {
	sugar.Fatalln(message)
}

func Panicf(message string, fields interface{}) {
	sugar.Panicf(message, fields)
}

func Panicln(message string) {
	sugar.Panicln(message)
}

func Panic(message string) {
	sugar.Panic(message)
}
