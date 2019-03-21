package log

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func init() {
	l, err := DefaultLogger()
	if err != nil {
		panic(err)
	}
	logger = l
}

// DefaultLogger returns the default makisu logger.
func DefaultLogger() (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	l, err := config.Build()
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}

// SetLogger sets the default logger.
func SetLogger(log *zap.SugaredLogger) {
	logger = log
}

// GetLogger returns the current SugaredLogger.
func GetLogger() *zap.SugaredLogger {
	return logger
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	GetLogger().Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

// With adds a variadic number of fields to the logging context.
// It accepts a mix of strongly-typed zapcore.Field objects and loosely-typed key-value pairs.
func With(args ...interface{}) *zap.SugaredLogger {
	return GetLogger().With(args...)
}