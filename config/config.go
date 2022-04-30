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

	Store store.Store
}

// New return a configure store with map stored and file stored.
// And it detects file format by file name extension.
// JSON and TOML are supported now. Otherwise, return error.
func New(dsn string) *Config {
	c := &Config{
		data:  make(map[string]interface{}),
		flat:  make(map[string]interface{}),
		env:   make(map[string]interface{}),
		arg:   make(map[string]interface{}),
		Store: store.NewStore(dsn),
	}
	return c
}

// Load config data from file. Returns old data and error.
func (c *Config) Load() (map[string]interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	out := cloneMap(c.data)

	data, err := c.Store.LoadStore()
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

	return c.Store.SaveStore(c.data)
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

// Clone a new config
func (c *Config) Clone() *Config {
	c.mu.Lock()
	defer c.mu.Unlock()

	return &Config{
		data: cloneMap(c.data),
		flat: flattenMap(c.data, ""),
		env:  cloneMap(c.env),
		arg:  cloneMap(c.arg),
	}
}
