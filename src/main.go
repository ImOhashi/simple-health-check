package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
