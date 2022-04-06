package storage

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// implements Storage interface
type SQLStorage struct {
	Watcher
	db               *gorm.DB
	sqlType          string
	connectionString string
	dialector        gorm.Dialector
}

func NewSqlStorage(u *url.URL) *SQLStorage {
	var connectionString string
	var dialect gorm.Dialector

	switch u.Scheme {
	case "postgresql":
		// handle `postgresql` as the scheme to be compatible with
		// standard uri style postgresql connection strings (i.e. like psql)
		u.Scheme = "postgres"
		fallthrough
	case "postgres":
		connectionString = pgconn(u)
		// Open does not actually open the connection, only create the Dialector object
		dialect = postgres.Open(connectionString)
	case "mysql":
		connectionString = mysqlconn(u)
		dialect = mysql.Open(connectionString)
	case "sqlite3":
		connectionString = sqlite3conn(u)
		dialect = sqlite.Open(connectionString)
	default:
		// unreachable because our storage backend factory
		// function (contracts.go) already checks the url scheme.
		logrus.Panicf("unknown sql storage backend %s", u.Scheme)
	}

	return &SQLStorage{
		Watcher:          nil,
		db:               nil,
		sqlType:          u.Scheme,
		connectionString: connectionString,
		dialector:        dialect,
	}
}

func pgconn(u *url.URL) string {
	password, _ := u.User.Password()
	decodedQuery, err := url.QueryUnescape(u.RawQuery)
	if err != nil {
		logrus.Warnf("failed to unescape connection string query parameters - they will be ignored")
		decodedQuery = ""
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s",
		u.Hostname(),
		u.Port(),
		u.User.Username(),
		password,
		strings.TrimLeft(u.Path, "/"),
		decodedQuery,
	)
}

func mysqlconn(u *url.URL) string {
	password, _ := u.User.Password()
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?%s",
		u.User.Username(),
		password,
		u.Host,
		strings.TrimLeft(u.Path, "/"),
		u.RawQuery,
	)
}

func sqlite3conn(u *url.URL) string {
	return filepath.Join(u.Host, u.Path)
}

func (s *SQLStorage) Open() error {
	db, err := gorm.Open(s.dialector, &gorm.Config{
		Logger: &GormLogger{},
	})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to connect to %s", s.sqlType))
	}
	s.db = db

	// Migrate the schemasqlType
	s.db.AutoMigrate(&Device{})

	if s.sqlType == "postgres" {
		watcher, err := NewPgWatcher(s.connectionString, db.Take(&Device{}).Statement.Table)
		if err != nil {
			return errors.Wrap(err, "failed to create pg watcher")
		}
		s.Watcher = watcher
	} else if s.sqlType == "mysql" || s.sqlType == "sqlite3" {
		db.Scopes()
		s.Watcher = NewGormWatcher(db, db.Take(&Device{}).Statement.Table)
	} else {
		s.Watcher = NewInProcessWatcher()
	}

	return nil
}

func (s *SQLStorage) Close() error {
	if s.db != nil {
		db, err := s.db.DB()
		if err != nil {
			return err
		}
		return db.Close()
	}
	return nil
}

func (s *SQLStorage) Save(device *Device) error {
	logrus.Debugf("saving device %s", key(device))
	if err := s.db.Save(&device).Error; err != nil {
		return errors.Wrapf(err, "failed to write device")
	}
	s.Watcher.EmitAdd(device)
	return nil
}

func (s *SQLStorage) List(username string) ([]*Device, error) {
	var err error
	devices := []*Device{}
	if username != "" {
		err = s.db.Where("owner = ?", username).Find(&devices).Error
	} else {
		err = s.db.Find(&devices).Error
	}

	logrus.Debugf("found %d device(s)", len(devices))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read devices from sql")
	}
	return devices, nil
}

func (s *SQLStorage) Get(owner string, name string) (*Device, error) {
	device := &Device{}
	if err := s.db.Where("owner = ? AND name = ?", owner, name).First(&device).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to read device")
	}
	return device, nil
}

func (s *SQLStorage) GetByPublicKey(publicKey string) (*Device, error) {
	device := &Device{}
	if err := s.db.Where("public_key = ?", publicKey).First(&device).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to read device")
	}
	return device, nil
}

func (s *SQLStorage) Delete(device *Device) error {
	if err := s.db.Delete(&device).Error; err != nil {
		return errors.Wrap(err, "failed to delete device file")
	}
	s.Watcher.EmitDelete(device)
	return nil
}
