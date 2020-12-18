package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"gopkg.in/yaml.v2"
)

type Servers struct {
	Servers []string `yaml:"servers"`
}

func getServerList() []string {
	data, err := ioutil.ReadFile("../pkg/config/config.yaml")
	servers := Servers{}
	err = yaml.Unmarshal(data, &servers)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return servers.Servers
}

func getTarget() (string, string) {
	servers := getServerList()
	index := rand.Intn(len(servers))
	target := servers[index]
	targetURL, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	return targetURL.String(), targetURL.Host
}

func Handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	defer r.Body.Close()
	targetURL, host := getTarget()
	request, err := http.NewRequest(r.Method, targetURL, r.Body)
	request.Header = r.Header.Clone()
	r.Header.Set("host", host)
	if err != nil {
		fmt.Println("Oops!")
		return
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	for k, v := range response.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)

}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8000", nil)
}
