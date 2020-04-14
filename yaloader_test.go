package main

import (
	"github.com/gpbbit/gb_go_web/yaloader"
	"os"
	"testing"
)

func TestFileLoad(t *testing.T) {
	link := "https://yadi.sk/i/xV3LGCEwax8RfA"
	path, err := yaloader.FileLoader(link)
	if err != nil{
		t.Error(err)
	}
	if !fileExists(path){
		t.Error(err)
	}
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}