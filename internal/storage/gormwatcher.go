package storage

import (
	"reflect"
	"runtime"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GormWatcher struct {
	*gorm.DB
	table string
}

func NewGormWatcher(db *gorm.DB, table string) *GormWatcher {
	logrus.Debug("creating gorm watcher")
	return &GormWatcher{
		DB:    db,
		table: table,
	}
}

func (w *GormWatcher) OnAdd(cb Callback) {
	name := runtime.FuncForPC(reflect.ValueOf(cb).Pointer()).Name()
	logrus.Debugf("OnAdd callback name %s", name)
	w.Callback().Create().Register(name, func(tx *gorm.DB) {
		w.emit(cb, tx)
	})
}

func (w *GormWatcher) OnDelete(cb Callback) {
	name := runtime.FuncForPC(reflect.ValueOf(cb).Pointer()).Name()
	logrus.Debugf("OnDelete callback name %s", name)
	w.Callback().Delete().Register(name, func(tx *gorm.DB) {
		w.emit(cb, tx)
	})
}

func (w *GormWatcher) OnReconnect(cb func()) {
	// noop because the watcher can't reconnect
}

func (w *GormWatcher) emit(cb Callback, tx *gorm.DB) {
	if tx.Statement.Table == w.table {
		d, _ := tx.Statement.Vars[0].(**Device)
		cb(*d)
	}
}

func (w *GormWatcher) EmitAdd(device *Device) {
	// noop because we rely on gorm callback
}

func (w *GormWatcher) EmitDelete(device *Device) {
	// noop because we rely on gorm callback
}
