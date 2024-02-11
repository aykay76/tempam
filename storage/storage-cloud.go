package storage

import "fmt"

type cloudStorage struct {
}

func (s *cloudStorage) StoreBlob(collectionName string, name string, content interface{}) error {
	fmt.Println("uploading to cloud")
	return nil
}

func (s *cloudStorage) ListBlobs(collectionName string, pattern string) ([]string, error) {
	return nil, nil
}

func (s *cloudStorage) GetBlob(collectionName string, name string) ([]byte, error) {
	return nil, nil
}

func (s *cloudStorage) GetAllBlobs(collectionName string, pattern string) ([][]byte, error) {
	return nil, nil
}

func (s *cloudStorage) DeleteBlob(collectionName string, name string) error {
	return nil
}

func CloudStorage() Storage {
	return &cloudStorage{}
}
