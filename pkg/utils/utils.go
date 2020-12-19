package utils

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Servers struct {
	Servers []string `yaml:"servers"`
}

func GetServerList() []string {
	data, err := ioutil.ReadFile(os.Getenv("config_address"))
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
