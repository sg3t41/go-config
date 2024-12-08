package main

type Endpoint struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type API struct {
	Port      int        `yaml:"port"`
	Debug     bool       `yaml:"debug"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Frontend struct {
	Port        int      `yaml:"port"`
	EnableHTTPS bool     `yaml:"enable_https"`
	Domains     []string `yaml:"domains"`
}

type Replica struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Database struct {
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	Username string    `yaml:"username"`
	Password string    `yaml:"password"`
	Replicas []Replica `yaml:"replicas"`
}

type Upstream struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Proxy struct {
	Host      string     `yaml:"host"`
	Port      int        `yaml:"port"`
	SSL       bool       `yaml:"ssl"`
	Upstreams []Upstream `yaml:"upstreams"`
}

type Scheme struct {
	API      API      `yaml:"api"`
	Frontend Frontend `yaml:"frontend"`
	Database Database `yaml:"database"`
	Proxy    Proxy    `yaml:"proxy"`
}
