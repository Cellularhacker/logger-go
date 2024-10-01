package logger

import (
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var L *zap.SugaredLogger

// Init -initializes the logger, different in test and production, call after config.Init
func Init(isProductionMode bool) {
	if isProductionMode {
		zc, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		L = zc.Sugar()
		return
	}

	zc := zap.NewDevelopmentEncoderConfig()
	zc.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zc.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	L = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zc),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	)).Sugar()
}

func Fatal(args ...interface{}) {
	if L == nil {
		log.Fatal(args)
		return
	}

	L.Fatal(args)
}
func Fatalf(format string, args ...interface{}) {
	if L == nil {
		log.Fatalf(format, args)
		return
	}

	L.Fatalf(format, args)
}
func Fatalln(args ...interface{}) {
	if L == nil {
		log.Fatalln(args)
		return
	}

	L.Fatalln(args)
}

func Info(args ...interface{}) {
	if L == nil {
		log.Info(args)
		return
	}

	L.Info(args)
}
func Infof(format string, args ...interface{}) {
	if L == nil {
		log.Infof(format, args)
		return
	}

	L.Infof(format, args)
}
func Infoln(args ...interface{}) {
	if L == nil {
		log.Fatalln(args)
		return
	}

	L.Infoln(args)
}

func Warn(args ...interface{}) {
	if L == nil {
		log.Warn(args)
		return
	}

	L.Warn(args)
}
func Warnf(format string, args ...interface{}) {
	if L == nil {
		log.Warnf(format, args)
		return
	}

	L.Warnf(format, args)
}
func Warnln(args ...interface{}) {
	if L == nil {
		log.Fatalln(args)
		return
	}

	L.Warnln(args)
}

func Debug(args ...interface{}) {
	if L == nil {
		log.Debug(args)
		return
	}

	L.Debug(args)
}
func Debugf(format string, args ...interface{}) {
	if L == nil {
		log.Debugf(format, args)
		return
	}

	L.Debugf(format, args)
}
func Debugln(args ...interface{}) {
	if L == nil {
		log.Fatalln(args)
		return
	}

	L.Debugln(args)
}

func Error(args ...interface{}) {
	if L == nil {
		log.Error(args)
		return
	}

	L.Error(args)
}
func Errorf(format string, args ...interface{}) {
	if L == nil {
		log.Errorf(format, args)
		return
	}

	L.Errorf(format, args)
}
func Errorln(args ...interface{}) {
	if L == nil {
		log.Fatalln(args)
		return
	}

	L.Errorln(args)
}
