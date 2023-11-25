package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"fakesite",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, site := range sites {
		checkSite(site)
	}

}

func checkSite(site string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is down!")
		return
	}

	fmt.Println(site, "is up!")
	// defer up.Body.Close()
}
