package main

import "confarp"

func main() {
	builder := confarp.ConfigurationBuilder{}
	builder.Add(confarp.NewMemoryConfigurationSource(map[string]string{
		"a": "1",
	}))
	configurationRoot := builder.Build()

	value := configurationRoot.Get("a")

	println(value)

}
