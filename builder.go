package confarp

import "confarp/abstraction"

type ConfigurationBuilder struct {
	sources    []abstraction.IConfigurationSource
	properties map[string]interface{}
}

func (c *ConfigurationBuilder) GetProperties() map[string]interface{} {
	return c.properties
}

func (c *ConfigurationBuilder) GetSources() []abstraction.IConfigurationSource {
	return c.sources
}

func (c *ConfigurationBuilder) Add(source abstraction.IConfigurationSource) abstraction.IConfigurationBuilder {
	c.sources = append(c.sources, source)
	return c
}

func (c *ConfigurationBuilder) Build() abstraction.IConfigurationRoot {
	providers := []abstraction.IConfigurationProvider{}
	for _, source := range c.sources {
		provider := source.Build(c)
		providers = append(providers, provider)
	}
	return NewConfigurationRoot(providers)
}
