package config

type Config struct {
	Proxy     Proxy               `mapstructure:"proxy"`
	Upstreams map[string]Upstream `mapstructure:"upstreams"`
}

type Proxy struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Upstream struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
