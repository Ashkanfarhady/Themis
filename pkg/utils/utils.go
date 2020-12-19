package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Servers struct {
	Servers []string `yaml:"servers"`
}

func GetServerList(config string) []string {
	data, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	servers := Servers{}
	err = yaml.Unmarshal(data, &servers)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return servers.Servers
}
