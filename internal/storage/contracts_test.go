package storage

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMemoryStorage(t *testing.T) {
	require := require.New(t)

	s, err := NewStorage("memory://")
	require.NoError(err)

	require.IsType(&InMemoryStorage{}, s)
}

func TestPostgresqlStorage(t *testing.T) {
	require := require.New(t)

	s, err := NewStorage("postgresql://localhost:5432/dbname?sslmode=disable")
	require.NoError(err)

	require.IsType(&SQLStorage{}, s)
}

func TestMysqlStorage(t *testing.T) {
	require := require.New(t)

	s, err := NewStorage("mysql://localhost:1234/dbname?sslmode=disable")
	require.NoError(err)

	require.IsType(&SQLStorage{}, s)
}

func TestSqliteStorage(t *testing.T) {
	require := require.New(t)

	s, err := NewStorage("sqlite3:///some/path/sqlite.db")
	require.NoError(err)

	require.IsType(&SQLStorage{}, s)
}

func TestSqliteStorageRelativePath(t *testing.T) {
	require := require.New(t)

	s, err := NewStorage("sqlite3://sqlite.db")
	require.NoError(err)

	require.IsType(&SQLStorage{}, s)
}

func TestUnknownStorage(t *testing.T) {
	require := require.New(t)

	s, err := NewStorage("foo://")
	require.Nil(s)
	require.Error(err)
	require.Equal(err.Error(), "unknown storage backend foo")
}

func TestCallbackOnAdd(t *testing.T) {
	require := require.New(t)

	s, err := NewStorage("memory://")
	require.NoError(err)
	err = s.Open()
	require.NoError(err)
	defer s.Close()

	testDevice := Device{}

	received := false
	s.OnAdd(func(device *Device) {
		received = true
	})

	err = s.Save(&testDevice)
	require.NoError(err)
	require.True(received, "OnAdd event received")
}

func TestCallbackOnAddSQLite(t *testing.T) {
	require := require.New(t)

	// Setup
	file, err := os.CreateTemp("", "testdb-*.sqlite3")
	require.NoError(err)
	defer file.Close()
	defer os.Remove(file.Name())

	s, err := NewStorage("sqlite3://")
	require.NoError(err)
	err = s.Open()
	require.NoError(err)
	defer s.Close()

	testDevice := Device{
		Name:      "001",
		Owner:     "admin",
		PublicKey: "61fa8c",
	}

	received := false
	s.OnAdd(func(device *Device) {
		require.Equal(*device, testDevice)
		received = true
	})

	// Act
	err = s.Save(&testDevice)
	require.NoError(err)
	require.Eventually(func() bool { return received }, time.Second, 10*time.Millisecond)

	// Assert
	require.True(received, "OnAdd event received")
}
