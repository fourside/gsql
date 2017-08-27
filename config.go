package gsql

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DbConfig DbConfig `yaml:"db"`
	SSH      SSH      `yaml:"ssh"`
}

type DbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"sshmode"`
}

type SSH struct {
	Forwarding bool `yaml:"forwarding"`
}

func ReadConfig() (Config, error) {
	var config Config
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
