package main

import (
	"net/http"
	"sync"
)

func checkURLs(urls []string) map[string]string {
	var wg sync.WaitGroup
	result := make(map[string]string)
	mu := sync.Mutex{}

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			resp, err := http.Get(u)
			status := "down"
			if err == nil && resp.StatusCode == 200 {
				status = "UP"
			}
			mu.Lock()
			result[u] = status
			mu.Unlock()
		}(url)
	}
	wg.Wait()
	return result

}
