package storage

import (
	"fmt"

	"github.com/ASlavchik/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	fd, ok := s.files[fileId]
	if !ok {
		return nil, fmt.Errorf("file %v is not exist", fd.ID)
	}
	return fd, nil
}
