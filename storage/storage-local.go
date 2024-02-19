package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type localStorage struct {
}

func LocalStorage() Storage {
	return &localStorage{}
}

func (s *localStorage) StoreBlob(collectionName string, name string, content interface{}) error {
	bytes, _ := json.Marshal(content)
	return os.WriteFile(name, bytes, 0644)
}

func (s *localStorage) ListBlobs(collectionName string, pattern string) ([]string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return nil, err
	}

	return matches, nil
}

func (s *localStorage) GetBlob(collectionName string, name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (s *localStorage) GetAllBlobs(collectionName string, filter interface{}, results interface{}) error {
	matches, err := filepath.Glob(filter.(string))
	if err != nil || len(matches) == 0 {
		return err
	}

	var blobs [][]byte
	blobs = results.([][]byte)

	for _, match := range matches {
		blob, err := os.ReadFile(match)
		if err != nil {
			return err
		}
		blobs = append(blobs, blob)
	}

	return nil
}

func (s *localStorage) DeleteBlob(collectionName string, name string) error {
	fmt.Println("Deleting blob: " + name)
	err := os.Remove(name)
	return err
}
