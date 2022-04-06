package storage

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// GormLogger is a custom logger for Gorm, making it use logrus.
type GormLogger struct {
}

func (l *GormLogger) LogMode(_ gormLog.LogLevel) gormLog.Interface {
	return l
}
func (*GormLogger) Info(_ context.Context, s string, v ...interface{}) {
	logrus.WithFields(
		logrus.Fields{
			"module":  "gorm",
			"type":    "logrus",
			"src_ref": utils.FileWithLineNum(),
		},
	).Infof(s, v...)
}
func (*GormLogger) Warn(_ context.Context, s string, v ...interface{}) {
	logrus.WithFields(
		logrus.Fields{
			"module":  "gorm",
			"type":    "logrus",
			"src_ref": utils.FileWithLineNum(),
		},
	).Warnf(s, v...)
}
func (*GormLogger) Error(_ context.Context, s string, v ...interface{}) {
	logrus.WithFields(
		logrus.Fields{
			"module":  "gorm",
			"type":    "logrus",
			"src_ref": utils.FileWithLineNum(),
		},
	).Errorf(s, v...)
}
func (*GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logrus.WithFields(
		logrus.Fields{
			"module":   "gorm",
			"type":     "sql",
			"rows":     rows,
			"src_ref":  utils.FileWithLineNum(),
			"duration": float64(elapsed.Nanoseconds()) / 1e6,
		},
	).Trace(sql)
}
