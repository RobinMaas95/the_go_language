// Find a web site that produces a large amount of data. Investigate caching by running fechall
// twice in succession to see wether the reported time changes much. Do you get the same content each
// time? Modify to print its output to a file so it can be examined.

// Answers:
// As we can see, the time decreases significantly the second time we call fetchall:
// > ./fetchall  http://9gag.com
//	0.91s   132617  http://9gag.com
//	0.91s elapsed
// > ./fetchall  http://9gag.com
//	0.21s   132595  http://9gag.com
//	0.21s elapsed

// At least in my test, the files do not contain the completely identical content:
// > cmp fetch_result_1649942814.txt fetch_result_1649942818.txt
//	fetch_result_1649942814.txt fetch_result_1649942818.txt differ: char 8709, line 60

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
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// Create a file for storing our response. Add current unix timestamp to be able to
	// compare multiple results later instead of overwritting our result each time
	out, err := os.Create(fmt.Sprintf("fetch_result_%d.txt", time.Now().Unix()))
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer out.Close()
	nbytes, err := io.Copy(out, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
