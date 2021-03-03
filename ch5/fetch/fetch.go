package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		filename, b, err := fetch( url )
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err )
			continue
		}
		fmt.Printf("%s %d\n", filename, b)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
