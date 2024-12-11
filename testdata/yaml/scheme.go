package testdatayaml

type Endpoint struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type API struct {
	Port      int        `yaml:"port"`
	Debug     bool       `yaml:"debug"`
	LogLevel  string     `yaml:"log_level"`
	Endpoints []Endpoint `yaml:"endpoints"`
	Security  Security   `yaml:"security"`
}

type Security struct {
	EnableAuth     bool     `yaml:"enable_auth"`
	AllowedOrigins []string `yaml:"allowed_origins"`
}

type Frontend struct {
	Port             int      `yaml:"port"`
	EnableHTTPS      bool     `yaml:"enable_https"`
	Domains          []string `yaml:"domains"`
	StaticAssetsPath string   `yaml:"static_assets_path"`
	Caching          Caching  `yaml:"caching"`
}

type Caching struct {
	Enabled bool  `yaml:"enabled"`
	MaxAge  int64 `yaml:"max_age"`
}

type Replica struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Pool struct {
	MaxConnections    int    `yaml:"max_connections"`
	MinConnections    int    `yaml:"min_connections"`
	ConnectionTimeout string `yaml:"connection_timeout"`
}

type Database struct {
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	Username string    `yaml:"username"`
	Password string    `yaml:"password"`
	Replicas []Replica `yaml:"replicas"`
	Pool     Pool      `yaml:"pool"`
}

type Cache struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Timeout  string `yaml:"timeout"`
}

type Upstream struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Proxy struct {
	Host      string     `yaml:"host"`
	Port      int        `yaml:"port"`
	SSL       bool       `yaml:"ssl"`
	Upstreams []Upstream `yaml:"upstreams"`
}

type Logging struct {
	Level    string   `yaml:"level"`
	Output   string   `yaml:"output"`
	Rotation Rotation `yaml:"rotation"`
}

type Rotation struct {
	MaxSize    string `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     string `yaml:"max_age"`
}

type Features struct {
	EnableNewUI       bool `yaml:"enable_new_ui"`
	EnableBetaFeature bool `yaml:"enable_beta_feature"`
}

type Email struct {
	SMTPServer  string   `yaml:"smtp_server"`
	Port        int      `yaml:"port"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	FromAddress string   `yaml:"from_address"`
	Recipients  []string `yaml:"recipients"`
}

type Scheme struct {
	API      API      `yaml:"api"`
	Frontend Frontend `yaml:"frontend"`
	Database Database `yaml:"database"`
	Cache    Cache    `yaml:"cache"`
	Proxy    Proxy    `yaml:"proxy"`
	Logging  Logging  `yaml:"logging"`
	Features Features `yaml:"features"`
	Email    Email    `yaml:"email"`
}
