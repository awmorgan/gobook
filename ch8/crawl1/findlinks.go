// Crawl1 crawls web links starting with the command-line arguments.
//
// This version quickly exhausts available file descriptors.
// due to excessive concurrent calls to links.Extract
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awmorgan/gobook/ch5/links"
)

func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() {
		worklist <- os.Args[1:]
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
