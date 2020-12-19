package main

import (
	"fmt"
	"os"

	"github.com/AshkanFarhady/Themis/pkg/handlers"
)

func main() {
	ConfigAddress := os.Getenv("config_address")
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}
	PORT := ":" + arguments[1]
	handlers.MainHandler(ConfigAddress, PORT)

}
