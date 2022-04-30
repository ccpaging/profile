package config

import (
	"os"
	"strings"
)

var DELIMITER string = "/"

// SetEnv sets the map of env. variable. It scans and filters
// all env. variable by prefix string and reference mas.
//  PREFIX_SECTION.KEY	SetEnv([]string{"PREFIX_"}, strings.NewReplacer(".", config.DELIMITER))
func (c *Config) SetEnv(prefixes []string, replacer *strings.Replacer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	out := make(map[string]interface{})
	env := os.Environ()
	for _, kv := range env {
		parts := strings.SplitN(kv, "=", 2)
		k, v := strings.ToUpper(parts[0]), parts[1]
		for _, f := range prefixes {
			if strings.HasPrefix(k, strings.ToUpper(f)) {
				k = strings.TrimPrefix(k, f)
				if replacer != nil {
					k = replacer.Replace(k)
				}
				println(k, v)
				out[k] = v
			}
		}
	}
	c.env = out
}

// Env returns the map of env. variable.
func (c *Config) Env() map[string]interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	return cloneMap(c.env)
}

// SetArg sets the map of arguments. It scans and filters
// all env. variable by reference mas.
func (c *Config) SetArg(in map[string]string, replacer *strings.Replacer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	out := make(map[string]interface{})
	for k, v := range in {
		k = strings.ToUpper(k)
		if replacer != nil {
			k = replacer.Replace(k)
		}
		out[k] = v
	}
	c.arg = out
}

// Arg returns the map of arguments.
func (c *Config) Arg() map[string]interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	return cloneMap(c.arg)
}
