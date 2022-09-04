package errorhandling

import (
	"fmt"
	"net/http"
	"sync"
)

func BasicErrorHandling() {

	visitURL := func(urls []string, wg *sync.WaitGroup) {
		for i := 0; i < len(urls); i++ {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(resp.StatusCode)
			}(urls[i])
		}
	}

	urls := []string{
		"https://www.google.com",
		"a",
		"b",
		"c",
		"d",
	}
	var wg sync.WaitGroup
	visitURL(urls, &wg)
	wg.Wait()
}

func AdvancedErrorHandling() {

	type response struct {
		resp *http.Response
		err  error
	}

	visitURL := func(urls []string, done <-chan interface{}) <-chan response {
		responses := make(chan response)

		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				select {
				case <-done:
					return
				case responses <- response{resp: resp, err: err}:
				}
			}
		}()

		return responses
	}

	urls := []string{
		"https://www.google.com",
		"a",
		"b",
		"c",
		"d",
	}

	errLimit := 2
	count := 0
	done := make(chan interface{})
	responses := visitURL(urls, done)
	for i := range responses {
		if i.err != nil {
			fmt.Println("error")
			count++
			if count == errLimit {
				fmt.Println("too many errors")
				close(done)
				break
			}
		} else {
			fmt.Println(i.resp.StatusCode)
		}

	}

}
