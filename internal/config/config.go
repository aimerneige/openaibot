package config

var Conf *Config

type CommonConfig struct {
	WSServer      string   `yaml:"ws_server"`
	WSToken       string   `yaml:"ws_token"`
	NickName      []string `yaml:"nickname"`
	CommandPrefix string   `yaml:"command_prefix"`
	SuperUsers    []int64  `yaml:"super_users"`
}

type OpenaiSecret struct {
	Account string `yaml:"account"`
	AppId   string `yaml:"appid"`
	Token   string `yaml:"token"`
	AesKey  string `yaml:"aeskey"`
}

type Config struct {
	Common CommonConfig `yaml:"common"`
	Secret OpenaiSecret `yaml:"secret"`
}
