package cache

type Cache interface {
	SetValue(key, value []byte) error
	GetValue(key []byte) ([]byte, error)
}
