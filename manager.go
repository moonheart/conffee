package confarp

import (
	"confarp/abstraction"
	"confarp/primitives"
)

type ConfigurationManager struct {
	sources                  *configurationSources
	properties               *configurationBuilderProperties
	providers                []abstraction.IConfigurationProvider
	changeToken              *primitives.ChangeToken
	changeTokenRegistrations []*primitives.ChangeToken
}

func NewConfigurationManager() *ConfigurationManager {
	c := new(ConfigurationManager)
	c.sources = newConfigurationSources(c)
	c.properties = newConfigurationBuilderProperties(c)
	c.changeToken = primitives.NewChangeToken()
	return c
}

func (c *ConfigurationManager) Get(key string) string {
	return getConfiguration(c.providers, key)
}

func (c *ConfigurationManager) Set(key, value string) {
	setConfiguration(c.providers, key, value)
}

func (c *ConfigurationManager) GetSection(key string) abstraction.IConfigurationSection {
	return NewConfigurationSection(c, key)
}

func (c *ConfigurationManager) GetChildren() []abstraction.IConfigurationSection {
	return GetChildrenImplementation(c, "")
}

func (c *ConfigurationManager) GetReloadChan() *primitives.ChangeToken {
	return c.changeToken
}

func (c *ConfigurationManager) Reload() {
	for _, provider := range c.providers {
		provider.Load()
	}
	c.raiseChanged()
}

func (c *ConfigurationManager) raiseChanged() {
	newChangeToken := primitives.NewChangeToken()
	c.changeToken, newChangeToken = newChangeToken, c.changeToken
	newChangeToken.OnReload()
}
func (c *ConfigurationManager) addSource(source abstraction.IConfigurationSource) {
	provider := source.Build(c)
	provider.Load()
	c.changeTokenRegistrations = append(c.changeTokenRegistrations, provider.GetReloadChan())
	c.providers = append(c.providers, provider)
	c.raiseChanged()
}

func (c *ConfigurationManager) reloadSorces() {

	var newTokens []*primitives.ChangeToken
	var newProvidersList []abstraction.IConfigurationProvider
	for _, source := range c.sources.sources {
		newProvidersList = append(newProvidersList, source.Build(c))
	}
	for _, provider := range newProvidersList {
		provider.Load()
		newTokens = append(newTokens, provider.GetReloadChan())
	}
	c.providers = newProvidersList
	c.raiseChanged()
}

func (c *ConfigurationManager) GetProviders() []abstraction.IConfigurationProvider {
	return c.providers
}

func (c *ConfigurationManager) GetProperties() map[string]interface{} {
	return c.properties.properties
}

func (c *ConfigurationManager) GetSources() []abstraction.IConfigurationSource {
	return c.sources.sources
}

func (c *ConfigurationManager) Add(source abstraction.IConfigurationSource) abstraction.IConfigurationBuilder {
	c.sources.sources = append(c.sources.sources, source)
	return c
}

func (c *ConfigurationManager) Build() abstraction.IConfigurationRoot {
	return c
}

type sources []abstraction.IConfigurationSource
type configurationSources struct {
	sources
	config *ConfigurationManager
}

func newConfigurationSources(config *ConfigurationManager) *configurationSources {
	return &configurationSources{
		config: config,
	}
}

func (cs *configurationSources) Get(index int) abstraction.IConfigurationSource {
	return cs.sources[index]
}
func (cs *configurationSources) Set(index int, value abstraction.IConfigurationSource) {
	cs.sources[index] = value
	cs.config.reloadSorces()
}

func (cs *configurationSources) Add(source abstraction.IConfigurationSource) {
	cs.sources = append(cs.sources, source)
	cs.config.addSource(source)
}

type properties map[string]interface{}
type configurationBuilderProperties struct {
	properties
	config *ConfigurationManager
}

func newConfigurationBuilderProperties(config *ConfigurationManager) *configurationBuilderProperties {
	return &configurationBuilderProperties{
		config: config,
	}
}
