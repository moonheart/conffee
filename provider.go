package confarp

import (
	"confarp/abstraction"
	"confarp/primitives"
	"sort"
	"strings"
)

type ConfigurationProvider struct {
	data        map[string]string
	reloadToken *primitives.ChangeToken
}

func NewConfigurationProvider() *ConfigurationProvider {
	return &ConfigurationProvider{
		data:        map[string]string{},
		reloadToken: primitives.NewChangeToken(),
	}
}

func (c *ConfigurationProvider) Get(key string) (string, bool) {
	if value, ok := c.data[key]; ok {
		return value, true
	}
	return "", false
}

func (c *ConfigurationProvider) Set(key, value string) {
	c.data[key] = value
}

func (c *ConfigurationProvider) GetReloadChan() *primitives.ChangeToken {
	return c.reloadToken
}

func (c *ConfigurationProvider) Load() {

}

func (c *ConfigurationProvider) GetChildKeys(earlierKeys []string, parentPath string) []string {
	results := []string{}

	if parentPath == "" {
		for k := range c.data {
			results = append(results, c.segment(k, 0))
		}
	} else {
		for k := range c.data {
			if len(k) > len(parentPath) &&
				strings.HasPrefix(k, parentPath) &&
				k[len(parentPath)] == ':' {
				results = append(results, c.segment(k, len(parentPath)+1))
			}
		}
	}
	results = append(results, earlierKeys...)
	sort.Strings(results)
	return results
}

func (c *ConfigurationProvider) segment(key string, prefixLength int) string {
	indexOf := strings.Index(key[prefixLength:], string(abstraction.KeyDelimiter))
	if indexOf < 0 {
		return key[prefixLength:]
	}
	return key[prefixLength:indexOf]
}
