package storage

type Storage interface {
	StoreBlob(name string, content []byte) error
	ListBlobs(pattern string) ([]string, error)
	GetAllBlobs(pattern string) ([][]byte, error)
	GetBlob(name string) ([]byte, error)
	DeleteBlob(name string) error
}
