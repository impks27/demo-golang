package main

import (
	"demo-golang/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
