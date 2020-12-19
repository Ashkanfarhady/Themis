package main

import (
	"os"

	"github.com/AshkanFarhady/Themis/pkg/handlers"
)

func main() {
	ConfigAddress := os.Getenv("config_address")

	handlers.MainHandler(ConfigAddress)

}
