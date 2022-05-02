package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type jsonFile struct {
	fileStore
}

func newJSONFile(path string) *jsonFile {
	f := new(jsonFile)
	f.path = path
	return f
}

func (f *jsonFile) Load() (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	out := make(map[string]interface{})
	if err := json.Unmarshal(content, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (f *jsonFile) Save(in map[string]interface{}) error {
	if in == nil {
		return errors.New("Config is nil")
	}
	content, err := json.MarshalIndent(in, "", "    ")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(f.path, content, 0644); err != nil {
		return err
	}
	return nil
}
