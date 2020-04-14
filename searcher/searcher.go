package searcher

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func SearchWordsOnPage(query string, sites []string) ([]string, error) {
	result := make([]string, 0, len(sites))
	var mux sync.Mutex
	if len(sites) == 0 {
		return result, errors.New("Sites count is 0")
	}
	var wg sync.WaitGroup
	for _, val := range sites {
		wg.Add(1)
		go Searching(val, query, &result, &wg, &mux)
	}
	wg.Wait()
	return result, nil
}

func Searching(site string, query string, result *[]string, wg *sync.WaitGroup, mux *sync.Mutex) {
	defer wg.Done()
	resp, err := http.Get(site)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bQuery := []byte(query)
	if bytes.Contains(body, bQuery) {
		mux.Lock()
		*result = append(*result, site)
		mux.Unlock()
	}
}
