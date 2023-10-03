package library

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func ModuleLog() LoggerZap {

	return LoggerZap{Logger: nil}
}

func Writelog(c echo.Context, env Env, logLevel string, msg string) {

	atomicLevel := zap.NewAtomicLevel()
	switch logLevel {
	case "debug":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case "err":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	}

	endoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	writer := &lumberjack.Logger{
		Filename:   env.LogOutput,
		MaxSize:    100, //100mb
		MaxAge:     30,
		MaxBackups: 5, //5compress
		LocalTime:  true,
		Compress:   true,
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(endoderConfig),
		zapcore.AddSync(writer),
		atomicLevel,
	)

	logger := zap.New(zapCore, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	logger.Log(atomicLevel.Level(), msg, zap.String("request_id", c.Response().Header().Get(echo.HeaderXRequestID)))

	defer logger.Sync()

}

// library.Writelog(c, u.env, "debug", "created")
