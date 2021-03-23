package conf

import "github.com/jasontconnell/conf"

type Config struct {
	Binding     string    `json:"binding"`
	BasePath    string    `json:"basePath"`
	ContentType string    `json:"contentType"`
	Mappings    []Mapping `json:"mappings"`
}

type Mapping struct {
	Source      string `json:"src"`
	Destination string `json:"dest"`
}

func LoadConfig(filename string) Config {
	cfg := Config{}

	conf.LoadConfig(filename, &cfg)

	return cfg
}
