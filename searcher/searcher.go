package searcher

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func SearchWordsOnPage(query string, sites []string) ([]string, error) {
	result := make([]string, 0, len(sites))
	if len(sites) == 0 {
		return result, errors.New("Sites count is 0")
	}

	for _, val := range sites {
		resp, err := http.Get(val)
		if err != nil {
			return result, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return result, err
		}
		bQuery := []byte(query)
		if bytes.Contains(body, bQuery) {
			result = append(result, val)
		}
		resp.Body.Close()
	}

	return result, nil
}
