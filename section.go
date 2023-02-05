package confarp

import (
	"confarp/abstraction"
	"confarp/primitives"
)

type ConfigurationSection struct {
	root abstraction.IConfigurationRoot
	path string
	key  string
}

func NewConfigurationSection(root abstraction.IConfigurationRoot, path string) *ConfigurationSection {
	return &ConfigurationSection{root, path, ""}
}

func (c *ConfigurationSection) Get(key string) string {
	return c.root.Get(abstraction.Combine(c.path, key))
}

func (c *ConfigurationSection) Set(key, value string) {
	c.root.Set(abstraction.Combine(c.path, key), value)
}

func (c *ConfigurationSection) GetSection(key string) abstraction.IConfigurationSection {
	return c.root.GetSection(abstraction.Combine(c.path, key))
}

func (c *ConfigurationSection) GetChildren() []abstraction.IConfigurationSection {
	return c.root.GetChildren()
}

func (c *ConfigurationSection) GetReloadChan() *primitives.ChangeToken {
	return c.root.GetReloadChan()
}

func (c *ConfigurationSection) GetKey() string {
	if c.key == "" {
		return abstraction.GetSectionKey(c.path)
	}
	return c.key
}

func (c *ConfigurationSection) GetPath() string {
	return c.path
}

func (c *ConfigurationSection) GetValue() string {
	return c.root.Get(c.path)
}

func (c *ConfigurationSection) SetValue(value string) {
	c.root.Set(c.path, value)
}
