package store

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type tomlFile struct {
	fileStore
}

func newTOMLFile(path string) *tomlFile {
	f := new(tomlFile)
	f.path = path
	return f
}

// load config file to config data only
func (f *tomlFile) LoadStore() (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	out := make(map[string]interface{})
	if err := toml.Unmarshal(content, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// save config data to file only
func (f *tomlFile) SaveStore(in map[string]interface{}) error {
	if in == nil {
		return errors.New("Config is nil")
	}

	var buf bytes.Buffer
	e := toml.NewEncoder(&buf)
	err := e.Encode(in)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(f.path, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}
