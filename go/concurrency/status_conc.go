package main

import (
	"fmt"
	"net/http"
)

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(site, "is up!")
	c <- "Yep it's up"
	// defer up.Body.Close()
}

func main() {
	sites := []string{
		"fakesite",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, site := range sites {
		// create a new thread - a goroutine
		go checkSite(site, c)
	}
	// checkSite(<-c, c)
	fmt.Println(<-c)
}

// Common go saying: Concurrency is not parallelism
// concurrency vs parallelism
// concurrency: doing many things at once (using one CPU)
// parallelism: doing many things at once, simultaneously (using multiple CPUs & cores)
