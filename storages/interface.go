package storages

type Storage interface {
	Get(string) (interface{}, error)
	// GetAll() ([]interface{}, error)
	// Flush() error
	Set(string, interface{}) error
}
