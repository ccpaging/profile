package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func setupStore(t *testing.T) (Store, func()) {
	tempDir, err := ioutil.TempDir("", "setupConfigFile")
	if err != nil {
		t.Fatal(err)
	}

	dsn := filepath.Join(tempDir, "config.toml")
	store := NewStore(dsn)
	return store, func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Fatal(err)
		}
	}
}

func TestStore(t *testing.T) {
	s, teardown := setupStore(t)
	defer teardown()

	t.Run("store", func(t *testing.T) {
		if s.String() == "" {
			t.Fatal("Config file name is not set")
		}
	})
}
