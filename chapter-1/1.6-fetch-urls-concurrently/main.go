/*
Fetches a get request through each URL concurrently.
The execution time of the program is equal only to
the longest request amongst all the fetches, instead
of being the sum of each fetch time

Example:
$ go run main.go http://gopl.io http://godoc.org http://youtube.com
> 0.54s     7490  http://godoc.org  200 OK
> 0.86s   567191  http://youtube.com  200 OK
> 1.01s     4154  http://gopl.io  200 OK
> 1.01s elapsed 

Note: Returns the html from the website, the GET response
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	var report string
	for range os.Args[1:] {
		s := <-ch
		fmt.Println(s)
		report += s + "\n"

	}

	s := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Printf(s)
	report += s + "\n"

	// Code below, This code appends a line of text to the file  
	// text.log. It creates the file if it doesn’t already exist.
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(report + "\n"); err != nil {
		log.Println(err)
	}

}

// fetch takes in a url string and a string channel, sends
// result of get request at the url to the string channel

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // sends to the channel ch
		return
	}

	// reads the body and copies it to the ioutil.Discard stream
	// Discard is an io.Writer on which all Write calls succeed without doing anything
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // Reads the body in one go
	resp.Body.Close()                                 // Don't leak resources

	if err != nil {
		ch <- fmt.Sprint(err) // sends to the channel ch
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s  %s", secs, nbytes, url, resp.Status)
}
