// Crawl2 crawls the web links starting with the command line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurren tcalls to links.Extract
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awmorgan/gobook/ch5/links"
)

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends in worklist
	// Start with command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ;n >0;n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func (lin string)  {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}
