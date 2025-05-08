// Name of the file is the domain from the URL.
// Content of the data fetched from a given URL will start with "-----Start-----\n"
// and end with "\n-----End-----\n".
// If the file exists data will be concatenated to it

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"regexp"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch (url, ch)	// start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)	// receive from channel ch
	}
	fmt.Printf("%.fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	file_name_pattern, err := regexp.Compile("[A-Za-z]+\\.[A-Za-z]+")
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	file_name := file_name_pattern.FindString(url)
	if file_name == "" {
		ch <- "Could not get a file name from URL"
		return
	}

	file_name += ".txt"
	// file, err := os.Create(file_name)
	file, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		file.Close()
		ch <- fmt.Sprint(err)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)	// send to channel ch
		return
	}

	file.WriteString("-----Start-----\n")
	nbytes, err := io.Copy(file, resp.Body)
	file.WriteString("\n-----End-----\n")
	resp.Body.Close()	// don't leak resources
	file.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}