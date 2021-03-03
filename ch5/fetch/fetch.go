package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		start1 := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// _, err = io.ReadAll(resp.Body)
		nbytes, err := io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s", b)
		fmt.Printf("%.2fs  %7d  %s\n", time.Since(start1).Seconds(), nbytes, url)
		
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
