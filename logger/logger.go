package logger

import (
	"context"
	"os"
	"time"

	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	// Logger data
	Logger struct {
		hostname string
		env      string
		zap      *zap.Logger
		ctx      context.Context
		request  interface{}
		response interface{}
		metadata metadata.Metadata
		timing   int64
		sugar    *zap.SugaredLogger
		level    zapcore.Level
	}

	// ILogger interface
	ILogger interface {
		Debug(mgs string)
		Info(mgs string)
		Warn(mgs string)
		Error(mgs string)
		Fatal(mgs string)
		Panic(mgs string)

		Debugf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Fatalf(format string, args ...interface{})
		Panicf(format string, args ...interface{})

		Debugw(mgs string, keysAndValues ...interface{})
		Infow(mgs string, keysAndValues ...interface{})
		Warnw(mgs string, keysAndValues ...interface{})
		Errorw(mgs string, keysAndValues ...interface{})
		Fatalw(mgs string, keysAndValues ...interface{})
		Panicw(mgs string, keysAndValues ...interface{})
	}
)

func NewLogger(hostname string, env string) (*Logger, error) {
	var zapLogger *zap.Logger
	var err error
	if env == "production" {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		zapLogger = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		))
	} else {
		zapLogger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	sugar := zapLogger.Sugar()

	return &Logger{
		hostname: hostname,
		env:      env,
		sugar:    sugar,
		zap:      zapLogger,
	}, nil
}

// LogWrapper service log wrapper
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// get metadata
		os, _ := os.Hostname()
		log, _ := NewLogger(os, "production")
		start := time.Now().UnixNano()
		log.level = zapcore.ErrorLevel
		log.metadata, _ = metadata.FromContext(ctx)

		// Before
		log.request = req.Body()

		// Run function handler
		err := fn(ctx, req, rsp)

		// After
		if err != nil {
			log.level = zapcore.ErrorLevel
		}
		log.timing = time.Now().UnixNano() - start
		log.response = rsp
		log.ctx = ctx

		// Add fields
		fields := []zap.Field{}

		fields = append(fields, zap.Field{
			Key:    "hostname",
			String: log.hostname,
			Type:   zapcore.StringType,
		})

		fields = append(fields, zap.Field{
			Key:     "timing",
			Integer: log.timing,
			Type:    zapcore.DurationType,
		})

		fields = append(fields, zap.Field{
			Key:       "response",
			Interface: log.response,
			Type:      zapcore.ReflectType,
		})

		fields = append(fields, zap.Field{
			Key:       "request",
			Interface: log.request,
			Type:      zapcore.ReflectType,
		})

		fields = append(fields, zap.Field{
			Key:       "metadata",
			Interface: log.metadata,
			Type:      zapcore.ReflectType,
		})

		fields = append(fields, zap.Field{
			Key:       "context",
			Interface: log.ctx,
			Type:      zapcore.ReflectType,
		})

		if err == nil {
			log.With(fields...).Info("success")
		} else {
			fields = append(fields, zap.Field{
				Key:       "error",
				Interface: err.Error(),
				Type:      zapcore.ReflectType,
			})

			log.With(fields...).Error("error")
		}

		return err
	}
}
