package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init initializes the logger and replaces the global zap logger with it
func Init(debug bool) {
	var consoleEncoder zapcore.EncoderConfig

	var level zapcore.Level

	if debug {
		consoleEncoder = zap.NewDevelopmentEncoderConfig()
		level = zap.DebugLevel
	} else {
		consoleEncoder = zap.NewProductionEncoderConfig()
		consoleEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
		level = zap.InfoLevel
	}

	consoleEncoder.EncodeLevel = nil

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoder),
		zapcore.AddSync(colorable.NewColorableStdout()),
		level,
	))
	defer logger.Sync()

	if debug {
		logger = logger.WithOptions(zap.AddCaller())
	}

	zap.ReplaceGlobals(logger)
}

// S is a shortcut for zap.S()
func S() *zap.SugaredLogger {
	return zap.S()
}
