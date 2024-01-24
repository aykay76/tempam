package storage

import "fmt"

type cloudStorage struct {
}

func (s *cloudStorage) StoreBlob(name string, content []byte) error {
	fmt.Println("uploading to cloud")
	return nil
}

func (s *cloudStorage) ListBlobs(pattern string) ([]string, error) {
	return nil, nil
}

func (s *cloudStorage) GetBlob(name string) ([]byte, error) {
	return nil, nil
}

func (s *cloudStorage) GetAllBlobs(pattern string) ([][]byte, error) {
	return nil, nil
}

func (s *cloudStorage) DeleteBlob(name string) error {
	return nil
}

func CloudStorage() Storage {
	return &cloudStorage{}
}
