package logger

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerOnce sync.Once
	logger     *zap.Logger
)

func Init() {
	loggerOnce.Do(func() {
		config := zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "@timestamp"
		config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
		logger, _ = config.Build(zap.AddCallerSkip(1))
	})
}
func Info(ctx context.Context, message string, messages map[string]interface{}) {
	if logger == nil {
		return
	}
	logger.Info(message, commonFields(ctx, message, messages)...)
}

func Error(ctx context.Context, message string, messages map[string]interface{}) {
	if logger == nil {
		return
	}
	logger.Error(message, commonFields(ctx, message, messages)...)
}

func commonFields(ctx context.Context, message string, messages map[string]interface{}) (fields []zap.Field) {
	requestId, ok := ctx.Value("X-Request-ID").(string)
	if ok {
		fields = append(fields, zap.String("request_id", requestId))
	}
	currentId, ok := ctx.Value("CurrentUserId").(uint)
	if ok {
		fields = append(fields, zap.Uint("actor_id", currentId))
	}
	for key, val := range messages {
		fields = append(fields, zap.Any(key, val))
	}

	return fields
}
