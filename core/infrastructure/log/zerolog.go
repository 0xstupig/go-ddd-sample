package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"io"
	"os"
)

type zeroLogger struct {
	logger *zerolog.Logger
}

func (l *zeroLogger) Error(args ...interface{}) {
	l.logger.Error().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Warn(args ...interface{}) {
	l.logger.Warn().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Info(args ...interface{}) {
	l.logger.Info().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Debug(args ...interface{}) {
	l.logger.Debug().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *zeroLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *zeroLogger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *zeroLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *zeroLogger) Errorw(msg string, args ...interface{}) {
	l.logger.Error().Fields(args).Msg(msg)
}

func (l *zeroLogger) Warnw(msg string, args ...interface{}) {
	l.logger.Warn().Fields(args).Msg(msg)
}

func (l *zeroLogger) Infow(msg string, args ...interface{}) {
	l.logger.Info().Fields(args).Msg(msg)
}

func (l *zeroLogger) Debugw(msg string, args ...interface{}) {
	l.logger.Debug().Fields(args).Msg(msg)
}

func NewLogger(cfg config.AppConfig) Logger {
	logCfg := cfg.Logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if v, err := zerolog.ParseLevel(logCfg.Level); err == nil {
		zerolog.SetGlobalLevel(v)
	}

	logWriter := func() io.Writer {
		if logCfg.Colorized {
			return io.MultiWriter(zerolog.NewConsoleWriter())
		}

		return io.MultiWriter(os.Stdout)
	}

	instanceLogger := zerolog.New(logWriter()).With().Timestamp().Caller().Logger()
	return &zeroLogger{
		logger: &instanceLogger,
	}
}
