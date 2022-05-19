package config

import (
	"sync"

	"github.com/ccpaging/profile/store"
)

type Config struct {
	mu   sync.Mutex
	data map[string]interface{}
	flat map[string]interface{}
	env  map[string]interface{}
	arg  map[string]interface{}

	file store.Store
}

// New return a configure store but not loaded.
func New(dsn string) *Config {
	return &Config{
		data: make(map[string]interface{}),
		flat: make(map[string]interface{}),
		env:  make(map[string]interface{}),
		arg:  make(map[string]interface{}),
		file: store.NewStore(dsn),
	}
}

// Load config data from file. 
// And it detects file format by file name extension. 
// JSON and TOML are supported now. Returns old map and error.
func (c *Config) Load() (map[string]interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	out := cloneMap(c.data)

	data, err := c.file.Load()
	if err != nil {
		return out, err
	}
	c.data = cloneMap(data)
	c.flat = flattenMap(c.data, "")
	return out, nil
}

// Save config data to file.
func (c *Config) Save() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.file.Save(c.data)
}

// Data returns the clone of config data.
func (c *Config) Data() map[string]interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	return cloneMap(c.data)
}

// Flatten returns the clone of config flatten map.
func (c *Config) Flatten() map[string]interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	return cloneMap(c.flat)
}
