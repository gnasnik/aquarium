package log

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _logger *Logger

var LEVELS = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"err":   zapcore.ErrorLevel,
}

type Logger struct {
	path  string
	rlog  *lumberjack.Logger // rolling logger
	log   *zap.Logger
	sugar *zap.SugaredLogger

	level zapcore.Level
	pid   []interface{}

	rolling        bool
	lastRotateTime time.Time
	lastRotateRW   sync.RWMutex
}

func NewLogger(path string, level string) (*Logger, error) {
	out := new(logger)
	out.rlog = new(lumberjack.Logger)

	out.path = path
	out.lastRotateTime = time.Now()
	out.level = LEVELS[level]
	out.pid = []interface{}{os.Getpid()}

	// config lumberjack
	out.rlog.Filename = path
	out.rlog.MaxSize = 0x1000 * 5 // automatic rolling file on it increment than 2GB
	out.rlog.LocalTime = true
	out.rlog.Compress = true
	out.rlog.MaxBackups = 60 // reserve last 60 day logs

	// config encoder config
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
	ec.EncodeTime = zapcore.ISO8601TimeEncoder

	// config core
	c := zapcore.AddSync(out.rlog)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(ec), c, out.level)
	out.log = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	).With(zap.Int("pid", os.Getpid()))

	out.rolling = true
	out.sugar = out.log.Sugar()
	return out, nil
}

func (tlog *Logger) checkRotate() {
	if !tlog.rolling {
		return
	}

	n := time.Now()

	tlog.lastRotateRW.Lock()
	defer tlog.lastRotateRW.Unlock()

	last := tlog.lastRotateTime
	y, m, d := last.Year(), last.Month(), last.Day()
	if y != n.Year() || m != n.Month() || d != n.Day() {
		go tlog.rlog.Rotate()
		tlog.lastRotateTime = n
	}

}

func (tlog *Logger) EnableDailyFile() {
	tlog.rolling = true
}

func (tlog *Logger) Err(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.ErrorLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Errorf(format, v...)
}

func (tlog *Logger) Errw(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.ErrorLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Errorw(format, v...)
}

func (tlog *Logger) Warn(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.WarnLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Warnf(format, v...)
}

func (tlog *Logger) Warnw(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.WarnLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Warnw(format, v...)
}

func (tlog *Logger) Info(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.InfoLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Infof(format, v...)
}

func (tlog *Logger) Infow(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.InfoLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Infow(format, v...)
}

func (tlog *Logger) Debug(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.DebugLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Debugf(format, v...)
}

func (tlog *Logger) Debugw(format string, v ...interface{}) {
	tlog.checkRotate()
	if !tlog.level.Enabled(zap.DebugLevel) {
		return
	}

	defer tlog.log.Sync()
	tlog.sugar.Debugw(format, v...)
}

func GetDefault() *Logger {
	return _logger
}

func SetDefault(l *Logger) {
	_logger = l
}

func Stdout() {
	l, _ := NewLogger("std", "debug")
	SetDefault(l)
}

func Err(format string, v ...interface{}) {
	_logger.Err(format, v...)
}

func Errw(format string, v ...interface{}) {
	_logger.Errw(format, v...)
}

func Info(format string, v ...interface{}) {
	_logger.Info(format, v...)
}

func Infow(format string, v ...interface{}) {
	_logger.Infow(format, v...)
}

func Debug(format string, v ...interface{}) {
	_logger.Debug(format, v...)
}

func Debugw(format string, v ...interface{}) {
	_logger.Debugw(format, v...)
}

func Warn(format string, v ...interface{}) {
	_logger.Warn(format, v...)
}

func Warnw(format string, v ...interface{}) {
	_logger.Warnw(format, v...)
}
