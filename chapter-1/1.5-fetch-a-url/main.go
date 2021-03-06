/*
Demonstrates the minimum necessary to retrieve
information over HTTP, fetches the content of each
specified URL

Example:
$ go run main.go http://gopl.io
Note: Returns the html from the website, the GET response
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting!")

	// iterates through each url given
	for _, url := range os.Args[1:] {

		// if url doesnt start with http:// or https:// add prexif
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			fmt.Println("added prefix")
			url = "http://" + url
		}

		// sends a get request to the url and returns the result in a response struct resp
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching: %v\n", err)
			os.Exit(1)
		}

		// ReadAll reads the entire response and stores result in b
		//? b in this case is a []byte containing the response
		//b, err := ioutil.ReadAll(resp.Body) //resp.Body is the server response as a readable stream
		//resp.Body.Close()                   // stream is closed to avoid leaking resources

		//! Alternative to ReadAll
		// io.Copy(dst, src) reads from src, writes to dst
		// writes to os.stdout directly and doesn't require a buffer large enough
		// to hold the entire stream
		//? b in this case is not a []byte but the amount of bytes copied instead
		b, err := io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v", url, err)
			os.Exit(1)
		}

		//? if using ioutil.ReadAll()
		//fmt.Printf("%s", b)

		//? if using io.Copy()
		fmt.Printf("%v\n", b)

		// Printing the response status:
		// Status     string // e.g. "200 OK"
    // StatusCode int    // e.g. 200
    // Proto      string // e.g. "HTTP/1.0"
    // ProtoMajor int    // e.g. 1
    // ProtoMinor int    // e.g. 0

		fmt.Printf("Response status code: %v\n", resp.Status)
		fmt.Printf("Response protocol: %v\n", resp.Proto)
	}
}
