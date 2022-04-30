package config

/*
	Accessing the sections is case-sensitive
*/

func (c *Config) GetSectionNames() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	var r []string

	for k, v := range c.data {
		if _, ok := v.(map[string]interface{}); ok {
			r = append(r, k)
		}
	}

	return r
}

func (c *Config) hasSection(section string) (map[string]interface{}, bool) {
	v, ok := c.data[section]
	if !ok {
		return nil, false
	}

	sectionMap, ok := v.(map[string]interface{})
	if !ok {
		return nil, false
	}
	return sectionMap, true
}

func (c *Config) HasSection(section string) (map[string]interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.hasSection(section)
}

func (c *Config) GetSection(section string) map[string]interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	sectionMap, ok := c.hasSection(section)
	if !ok {
		sectionMap = make(map[string]interface{})
		c.data[section] = sectionMap
	}
	return cloneMap(sectionMap)
}

func (c *Config) WriteSection(section string, m map[string]interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[section] = cloneMap(m)
	c.flat = flattenMap(c.data, "")
}

/*
	Accessing the keys is not case-sensitive
*/

func (c *Config) HasKey(section string, key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.flat[flattenKey(section, key)]
	return v, ok
}

func (c *Config) GetValue(section string, key string, i interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		v       interface{} = indirect(i)
		flatKey string      = flattenKey(section, key)
	)

	if newVal, ok := c.flat[flatKey]; ok {
		v = newVal
	} else {
		c.writeValue(section, key, i)
	}

	if newVal, ok := c.env[flatKey]; ok {
		v = newVal
	}

	if newVal, ok := c.arg[flatKey]; ok {
		v = newVal
	}

	return v
}

func (c *Config) writeValue(section string, key string, i interface{}) {
	i = indirect(i)

	sectionMap, ok := c.hasSection(section)
	if !ok {
		sectionMap = make(map[string]interface{})
		c.data[section] = sectionMap
	}
	sectionMap[key] = i

	c.flat[flattenKey(section, key)] = i
}

func (c *Config) WriteValue(section string, key string, i interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.writeValue(section, key, i)
}
