package store

import (
	"github.com/pkg/errors"
)

const dir = "files"

type FileStore struct {
	dir string
}

func NewFileStore(cache Cache) (*FileStore, error) {
	folder, err := cache.GetFile(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get file store folder: %v", dir)
	}

	return &FileStore{
		dir: folder,
	}, nil
}
