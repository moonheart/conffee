package confarp

import "confarp/abstraction"

type MemoryConfigurationSource struct {
	InitialData map[string]string
}

func NewMemoryConfigurationSource(initialData map[string]string) *MemoryConfigurationSource {
	return &MemoryConfigurationSource{InitialData: initialData}
}

func (m *MemoryConfigurationSource) Build(builder abstraction.IConfigurationBuilder) abstraction.IConfigurationProvider {
	return NewMemoryConfigurationProvider(m)
}

type MemoryConfigurationProvider struct {
	ConfigurationProvider
	source *MemoryConfigurationSource
}

func NewMemoryConfigurationProvider(source *MemoryConfigurationSource) *MemoryConfigurationProvider {
	provider := &MemoryConfigurationProvider{
		source: source,
	}
	provider.data = map[string]string{}
	for k, v := range source.InitialData {
		provider.data[k] = v
	}
	return provider
}

func (p *MemoryConfigurationProvider) Add(k, v string) {
	p.data[k] = v
}
