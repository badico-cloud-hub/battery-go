package storages

type StorageGet interface {
	Get(string) (interface{}, error)
}
type Storage interface {
	StorageGet
	// GetAll() ([]interface{}, error)
	// Flush() error
	Set(string, interface{}) error
}
