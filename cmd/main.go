package main

import (
	"net/http"

	"github.com/AshkanFarhady/Themis/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.ListenAndServe(":8000", nil)
}
