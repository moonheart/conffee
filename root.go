package confarp

import (
	"confarp/abstraction"
	"confarp/primitives"
)

type ConfigurationRoot struct {
	providers    []abstraction.IConfigurationProvider
	changeTokens []*primitives.ChangeToken
	changeChan   chan int
}

func NewConfigurationRoot(providers []abstraction.IConfigurationProvider) *ConfigurationRoot {
	root := ConfigurationRoot{providers: providers}
	for _, provider := range providers {
		provider.Load()
		root.changeTokens = append(root.changeTokens, provider.GetReloadChan())
	}
	return &root
}

func (c *ConfigurationRoot) Get(key string) string {
	return getConfiguration(c.providers, key)
}

func getConfiguration(providers []abstraction.IConfigurationProvider, key string) string {
	for i := len(providers) - 1; i >= 0; i-- {
		provider := providers[i]
		if value, ok := provider.Get(key); ok {
			return value
		}
	}
	return ""
}

func setConfiguration(providers []abstraction.IConfigurationProvider, key string, value string) {
	if len(providers) == 0 {
		return
	}
	for _, provider := range providers {
		provider.Set(key, value)
	}
}

func (c *ConfigurationRoot) Set(key, value string) {
	setConfiguration(c.providers, key, value)
}

func (c *ConfigurationRoot) GetSection(key string) abstraction.IConfigurationSection {
	return NewConfigurationSection(c, key)
}

func (c *ConfigurationRoot) GetChildren() []abstraction.IConfigurationSection {
	//TODO implement me
	panic("implement me")
}

func (c *ConfigurationRoot) GetReloadChan() *primitives.ChangeToken {
	//TODO implement me
	panic("implement me")
}

func (c *ConfigurationRoot) Reload() {
	//TODO implement me
	panic("implement me")
}

func (c *ConfigurationRoot) GetProviders() []abstraction.IConfigurationProvider {
	//TODO implement me
	panic("implement me")
}

func GetChildrenImplementation(root abstraction.IConfigurationRoot, path string) []abstraction.IConfigurationSection {
	var providers []abstraction.IConfigurationProvider
	if manager, ok := root.(*ConfigurationManager); ok {
		providers = manager.GetProviders()
	} else {
		providers = root.GetProviders()
	}

	keys := []string{}
	for _, provider := range providers {
		keys = provider.GetChildKeys(keys, path)
	}
	distinctKeys := map[string]struct{}{}
	for _, key := range keys {
		distinctKeys[key] = struct{}{}
	}

	children := []abstraction.IConfigurationSection{}
	for key := range distinctKeys {
		key := key
		if path != "" {
			key = abstraction.Combine(path, key)
		}
		children = append(children, root.GetSection(key))
	}
	return children
}
