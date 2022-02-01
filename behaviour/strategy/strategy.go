package strategy

import (
	"fmt"
)

// 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}
	return s, nil
}

type fileStorage struct{}

func (s *fileStorage) Save(name string, data []byte) error {
	fmt.Println(name)
	return nil
}

type encryptFileStorage struct{}

func (s *encryptFileStorage) Save(name string, data []byte) error {
	name = specila(name)
	fmt.Println(name)
	return nil
}

func specila(name string) string {
	name = name + name
	return name
}
