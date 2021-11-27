package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	//"log"
	//"net/http"
)

func handler_func(wr http.ResponseWriter, req *http.Request) {
	var urls = []string{
		"http://google.com",
		"http://yahoo.com",
		"http://somebadaddress_ithink.com",
		"http://msn.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		fmt.Println("spawing:", url)

		fmt.Fprintf(wr, "<html><body>")
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(wr, "%s %s<br/>", url, err.Error())
			} else {
				fmt.Fprintf(wr, "%s %d<br/>", url, resp.StatusCode)
			}
		}(url)

		wg.Wait()
		fmt.Fprintf(wr, "</body></html>")
	}

}

func main() {
	http.HandleFunc("/", handler_func)

	go func() {
		err := http.ListenAndServe(":8020", nil)
		log.Fatal(err)
	}()

	fmt.Print("Press enter to exit")
	fmt.Scanln()

	fmt.Print("Bye, exiting")
}
