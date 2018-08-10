package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	log.Println("waitGroup-0:")
	waitGroup0()
	log.Println("waitGroup-1:")
	waitGroup1()
}

func waitGroup0() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.sogou.com/",
		"http://weibo.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp, _ := http.Get(url)
			log.Println(url, resp.StatusCode)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func waitGroup1() {
	var num int
	var data interface{}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go getNumAsync(wg, &num)
	wg.Add(1)
	go getDataAsync(wg, &data)
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	log.Println("waitGroup1", num, data)
}

func getNumAsync(wg *sync.WaitGroup, num *int) {
	defer wg.Done()
	log.Println("getNumAsync---")
	*num = 3
	return
}
func getDataAsync(wg *sync.WaitGroup, data *interface{}) {
	defer wg.Done()
	log.Println("getDataAsync---")
	*data = map[string]interface{}{
		"name": "goith",
		"age":  "30",
	}
	return
}
