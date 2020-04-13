package main

import (
	"github.com/gpbbit/gb_go_web/searcher"
	"testing"
)

func TestSearchWordsOnPage(t *testing.T) {
	list := []string{
		"https://www.avtomaxi.ru/",
		"https://stackoverflow.com/",
	}
	sample := []string{
		"https://www.avtomaxi.ru/",
	}
	got, err := searcher.SearchWordsOnPage("прокат авто", list)
	if err != nil {
		t.Errorf("Error in method %v", err)
	}

	if len(got) != len(sample) {
		t.Errorf("incorrect length for got expected:%d, got:%d", len(sample), len(got))
	}
	for i, val := range got {
		if val != sample[i] {
			t.Errorf("Bad result for %v == %v , at index %d", val, sample[i], i)
		}
	}
}
func TestSearchWordsOnPageFOrZeroLength(t *testing.T) {
	list := make([]string, 0)
	_, err := searcher.SearchWordsOnPage("прокат авто", list)
	t.Logf("Got error: %v", err)
	if err == nil {
		t.Errorf("Expected error but got err = %v", err)
	}
}
