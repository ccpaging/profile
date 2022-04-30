package store

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// fileStore is a config store backed by a file such as config/config.json.
//
// It also uses the folder containing the configuration file for storing other configuration files.
// Not to be used directly. Only to be used as a backing store for config.Store
type fileStore struct {
	path string
}

func newFileStore(fullpath string) Store {
	if fullpath == "" {
		fp := os.Args[0]
		fplen, extlen := len(fp), len(filepath.Ext(fp))
		fullpath = fp[:(fplen-extlen)] + ".ini"
	}
	ext := strings.ToLower(filepath.Ext(fullpath))
	switch ext {
	case ".json":
		return newJSONFile(fullpath)
	case ".toml":
	case ".ini":
	default:
	}
	return newTOMLFile(fullpath)
}

// Existed returns true if the given file was previously persisted.
func (fs *fileStore) Existed() (bool, error) {
	if fs.path == "" {
		return false, nil
	}

	if _, err := os.Stat(fs.path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, errors.New("failed to check if file exists. " + err.Error())
		}
	}

	return true, nil
}

// Remove removes a previously persisted configuration file.
func (fs *fileStore) Remove() error {
	if fs.path == "" {
		return nil
	}

	err := os.Remove(fs.path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return errors.New("failed to remove file. " + err.Error())
	}
	return nil
}

func (fs *fileStore) String() string {
	return fs.path
}

func (fs *fileStore) Close() error {
	return nil
}
