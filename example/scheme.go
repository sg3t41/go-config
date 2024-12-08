// scheme.go
package main

type App1 struct {
	Key1 string `yaml:"key1"`
	Key2 string `yaml:"key2"`
}

type App2 struct {
	Key3 []string `yaml:"key3"`
}

type App3 struct {
	Key4 string `yaml:"key4"`
}

type Scheme struct {
	App1  App1    `yaml:"app1"`
	App2  App2    `yaml:"app2"`
	App3  App3    `yaml:"app3"`
	Data1 *string `yaml:"data1"`
	Data2 *string `yaml:"data2"`
	Data3 int     `yaml:"data3"`
}

