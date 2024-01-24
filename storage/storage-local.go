package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

type localStorage struct {
}

func LocalStorage() Storage {
	return &localStorage{}
}

func (s *localStorage) StoreBlob(name string, content []byte) error {
	return os.WriteFile(name, content, 0644)
}

func (s *localStorage) ListBlobs(pattern string) ([]string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return nil, err
	}

	return matches, nil
}

func (s *localStorage) GetBlob(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (s *localStorage) GetAllBlobs(pattern string) ([][]byte, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return nil, err
	}

	var blobs [][]byte

	for _, match := range matches {
		blob, err := os.ReadFile(match)
		if err != nil {
			return nil, err
		}
		blobs = append(blobs, blob)
	}

	return blobs, nil
}

func (s *localStorage) DeleteBlob(name string) error {
	fmt.Println("Deleting blob: " + name)
	err := os.Remove(name)
	return err
}
