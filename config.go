package main

import (
	"github.com/spf13/viper"
	"sync"
)

var Config *config
var once sync.Once

type config struct {
	IpPort       string `yaml:"ip_port"`
	TopicName    string `yaml:"topic_name"`
	Partitions   int    `yaml:"partitions"`
	Replications int    `yaml:"replications"`
}

func LoadConfig() *config {
	once.Do(func() {
		Config = &config{
			IpPort:       viper.GetString("ip_port"),
			TopicName:    viper.GetString("topic_name"),
			Partitions:   viper.GetInt("partitions"),
			Replications: viper.GetInt("replications"),
		}
	})
	return Config
}
