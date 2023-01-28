package abstraction

import (
	"container/list"
	"errors"
)

type IConfiguration interface {
	Get(key string) string
	Set(key, value string)
	GetSection(key string) IConfigurationSection
	GetChildren() []IConfigurationSection
	GetReloadChan() chan int
}

type IConfigurationSection interface {
	IConfiguration
	GetKey() string
	GetPath() string
	GetValue() string
	SetValue(value string)
}

type IConfigurationRoot interface {
	IConfiguration
	Reload()
	GetProviders() []*IConfigurationProvider
}

type IConfigurationProvider interface {
	Get(key string) (string, bool)
	Set(key, value string)
	GetReloadChan() chan int
	Load()
	GetChildKeys(earlierKeys []string, parentPath string)
}

type IConfigurationBuilder interface {
	GetProperties() map[string]interface{}
	GetSources() []*IConfigurationSource
	Add(source IConfigurationSource) *IConfigurationBuilder
	Build() *IConfigurationRoot
}

func Add[TSource IConfigurationSource](builder IConfigurationBuilder, configureSource func(*TSource)) *IConfigurationBuilder {
	var source TSource
	if configureSource != nil {
		configureSource(&source)
	}
	return builder.Add(source)
}

func GetConnectionString(configuration IConfiguration, name string) string {
	return configuration.GetSection("ConnectionStrings").Get(name)
}

func AsMap(configuration IConfiguration, makePathsRelative bool) (m map[string]string) {
	stack := list.List{}
	stack.PushBack(configuration)
	prefixLength := 0
	if rootSection, ok := configuration.(IConfigurationSection); ok {
		prefixLength = len(rootSection.GetPath())
	}
	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		config := element.Value.(IConfiguration)
		if section, ok := config.(IConfigurationSection); ok && (!makePathsRelative || section != configuration) {
			m[section.GetPath()[prefixLength:]] = section.GetValue()
		}
		for _, child := range config.GetChildren() {
			stack.PushBack(child)
		}
	}
	return
}

// Exists determines whether the section has a Value or has children
// section – The section to enumerate.
// returns true if the section has values or children; otherwise, false.
func Exists(section IConfigurationSection) bool {
	return section.GetValue() != "" || len(section.GetChildren()) > 0
}

func GetRequiredSection(configuration IConfiguration, key string) (IConfigurationSection, error) {
	if configuration == nil {
		return nil, errors.New("configuration is nil")
	}
	section := configuration.GetSection(key)
	if Exists(section) {
		return section, nil
	}
	return nil, errors.New("invalid section name")
}

type IConfigurationSource interface {
	Build(builder *IConfigurationBuilder)
}
