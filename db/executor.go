package db

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/swanwish/go-helper/logs"
)

const (
	ErrorMessage_GetConnectionFailed = "Failed to get database connection"
	ErrorMessageNoConnectionProvider = "Connection provider not specified"
)

var (
	ErrNoConnectionProvider = errors.New(ErrorMessageNoConnectionProvider)
)

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
}

type DefaultDB struct {
	ConnectionGetterFunc func() (*sqlx.DB, error)
}

func (d DefaultDB) Select(dest interface{}, query string, args ...interface{}) error {
	if d.ConnectionGetterFunc == nil {
		logs.Error("connection getter not specified.")
		return ErrNoConnectionProvider
	}
	dbConnection, err := d.ConnectionGetterFunc()
	if err != nil {
		logs.Error(ErrorMessage_GetConnectionFailed, err)
		return err
	}
	defer dbConnection.Close()

	return dbConnection.Select(dest, query, args...)
}

func (d DefaultDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.ConnectionGetterFunc == nil {
		logs.Error("connection provider not specified.")
		return nil, ErrNoConnectionProvider
	}
	dbConnection, err := d.ConnectionGetterFunc()
	if err != nil {
		logs.Error(ErrorMessage_GetConnectionFailed, err)
		return nil, err
	}
	defer dbConnection.Close()

	return dbConnection.Exec(query, args...)
}

func (d DefaultDB) Get(dest interface{}, query string, args ...interface{}) error {
	if d.ConnectionGetterFunc == nil {
		logs.Error("connection provider not specified.")
		return ErrNoConnectionProvider
	}
	dbConnection, err := d.ConnectionGetterFunc()
	if err != nil {
		logs.Error(ErrorMessage_GetConnectionFailed, err)
		return err
	}
	defer dbConnection.Close()

	return dbConnection.Get(dest, query, args...)
}
