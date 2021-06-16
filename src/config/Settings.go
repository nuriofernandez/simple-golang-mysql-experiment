package config

type Settings struct {
	Database string `yaml:"database"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`

	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}
