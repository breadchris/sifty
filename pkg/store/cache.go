package store

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

type Cache interface {
	GetFile(name string) (string, error)
	GetFolder(name string) (string, error)
}

type FolderCache struct {
	dir string
}

func createCacheFolder(name string) error {
	// Get the current user
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("could not get current user: %v", err)
	}

	// Create the .sifty folder in the user's home directory
	path := u.HomeDir + "/." + name
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0700); err != nil {
			return fmt.Errorf("could not create .%s folder: %v", name, err)
		}
	}
	return nil
}

func (c *FolderCache) GetFile(name string) (string, error) {
	return path.Join(c.dir, name), nil
}

func (c *FolderCache) GetFolder(name string) (string, error) {
	return path.Join(c.dir, name), nil
}

func NewFolderCache(name string) (*FolderCache, error) {
	if err := createCacheFolder(name); err != nil {
		return nil, err
	}

	return &FolderCache{
		dir: name,
	}, nil
}
