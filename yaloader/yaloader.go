package yaloader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type FData struct {
	Href      string
	Method    string
	Templated string
	Error     string
}

const APIURL string = "https://cloud-api.yandex.net/v1/disk/public/resources/download"
const SaveDir string = "download/"

func FileLoader(link string) (path string,err error) {
	downloadURL, err := APIRequest(link)
	checkError(err)
	uri, err := url.Parse(downloadURL)
	checkError(err)
	values := uri.Query()
	name := values.Get("filename")
	fileResp, err := http.Get(downloadURL)
	defer fileResp.Body.Close()
	checkError(err)
	fileBody, err := ioutil.ReadAll(fileResp.Body)
	checkError(err)
	path = SaveDir + name
	err = ioutil.WriteFile(path, fileBody, 0644)
	return
}

func APIRequest(link string) (downloadLink string, err error) {
	uri := APIURL + "?public_key=" + link
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var fileData FData
	json.Unmarshal(body, &fileData)
	switch fileData.Error {
	case "":
		downloadLink = fileData.Href
		err = nil
	case "TooManyRequestsError":
		downloadLink, err = APIRequest(link)
	default:
		downloadLink = ""
		err = errors.New("Disk Error: " + fileData.Error)
	}
	return
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
