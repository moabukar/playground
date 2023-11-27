package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is down!")
		c <- site
		return
	}

	fmt.Println(site, "is up!")
	c <- site
	// defer up.Body.Close()
}

func main() {
	sites := []string{
		"http://google.com",
		"http://twitter.com",
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
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	//////////////////////////////
	// Method 1
	// for {
	// 	//fmt.Println(<-c)
	// 	go checkSite(<-c, c)
	// }
	// Method 2 (same as Method 1 but more readable)
	for l := range c {
		//fmt.Println(<-c)
		// time.Sleep(5 * time.Second)
		go func(site string) {
			time.Sleep(5 * time.Second)
			checkSite(site, c)
		}(l)
	}
}

// Common go saying: Concurrency is not parallelism
// concurrency vs parallelism
// concurrency: doing many things at once (using one CPU)
// parallelism: doing many things at once, simultaneously (using multiple CPUs & cores)
