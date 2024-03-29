package storage

type Storage interface {
	StoreBlob(collectionName string, blobName string, content interface{}) error
	ListBlobs(collectionName string, pattern string) ([]string, error)
	GetAllBlobs(collectionName string, filter interface{}, results interface{}) error
	GetBlob(collectionName string, blobName string) ([]byte, error)
	DeleteBlob(collectionName string, blobName string) error
}
