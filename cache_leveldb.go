package cache

import (
	"errors"
	"ytserver/logs"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ErrParameterMissing = errors.New("Parameter is missing")
)

type CacheLevelDB struct {
}

func GetLevelDB() (*leveldb.DB, error) {
	return leveldb.OpenFile("leveldb", nil)
}

func (cache CacheLevelDB) SetValue(key, value []byte) error {
	if len(key) == 0 {
		logs.Error("The key is empty.")
		return ErrParameterMissing
	}
	db, err := GetLevelDB()
	if err != nil {
		logs.Errorf("Faieled to get level db connection, the error is %v\n", err)
		return err
	}
	defer db.Close()

	return db.Put(key, value, nil)
}

func (cache CacheLevelDB) GetValue(key []byte) ([]byte, error) {
	db, err := GetLevelDB()
	if err != nil {
		logs.Errorf("Failed to get level db connection, the error is %v\n", err)
		return []byte{}, err
	}
	defer db.Close()

	return db.Get(key, nil)
}
