package handlers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/AshkanFarhady/Themis/pkg/utils"
)

var configAddress string

func getTarget(servers []string) (string, string) {
	index := rand.Intn(len(servers))
	target := servers[index]
	targetURL, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	return targetURL.String(), targetURL.Host
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	ServerList := utils.GetServerList(configAddress)

	client := &http.Client{}
	defer r.Body.Close()
	targetURL, host := getTarget(ServerList)
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

func MainHandler(config string) {
	configAddress = config
	http.HandleFunc("/", LoadHandler)
	http.ListenAndServe(":8000", nil)

}
