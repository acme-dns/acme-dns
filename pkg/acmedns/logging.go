package acmedns

import (
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func SetupLogging(config AcmeDnsConfig) (*zap.Logger, error) {
	var (
		logger *zap.Logger
		zapCfg zap.Config
		err    error
	)

	logformat := "console"
	if config.Logconfig.Format == "json" {
		logformat = "json"
	}
	outputPath := "stdout"
	if config.Logconfig.Logtype == "file" {
		outputPath = config.Logconfig.File
	}
	errorPath := "stderr"
	if config.Logconfig.Logtype == "file" {
		errorPath = config.Logconfig.File
	}

	zapCfg.Level, err = zap.ParseAtomicLevel(config.Logconfig.Level)
	if err != nil {
		return logger, err
	}
	zapCfg.Encoding = logformat
	zapCfg.OutputPaths = []string{outputPath}
	zapCfg.ErrorOutputPaths = []string{errorPath}
	zapCfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:     "time",
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
	}

	logger, err = zapCfg.Build()
	return logger, err
}
