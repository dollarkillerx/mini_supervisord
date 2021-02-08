package config

type Servers struct {
	Servers map[string]ServersConfigItem `yaml:"Servers"`
}

type ServersConfigItem struct {
	Command string `yaml:"Command"`
	LogFile string `yaml:"LogFile"`
}
