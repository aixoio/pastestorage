package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/aixoio/pastestorage/converter"
)

type links_list struct {
	Index int
	Link  string
}

type final_dat struct {
	Filename string
	Links    []links_list
}

func UploadFile(filename, api_key, username, password string) {
	pastes, err := converter.ConvertFileToText(filename)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	user_key, _ := loginPost(username, password, api_key)
	links := []links_list{}
	var links_mutex sync.Mutex

	for i, paste := range pastes {
		wg.Add(1)
		go func(paste string, i int) {
			link, _ := postPaste(paste, api_key, user_key)
			links_mutex.Lock()
			links = append(links, links_list{
				Index: i,
				Link:  link,
			})
			links_mutex.Unlock()
			wg.Done()
		}(paste, i)
	}

	wg.Wait()

	last_paste := final_dat{
		Links:    links,
		Filename: filename,
	}

	json, _ := json.Marshal(last_paste)

	link, _ := postPaste(string(json), api_key, user_key)

	fmt.Println(link)

}

func loginPost(username, password, api_key string) (string, error) {
	api_endpoint := "https://pastebin.com/api/api_login.php"

	request_dat := url.Values{}
	request_dat.Set("api_dev_key", api_key)
	request_dat.Set("api_user_name", username)
	request_dat.Set("api_user_password", password)

	req, err := http.NewRequest("POST", api_endpoint, strings.NewReader(request_dat.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf := new(strings.Builder)
	io.Copy(buf, res.Body)

	if res.StatusCode != 200 {
		return "", errors.New("status code is not 200")
	}

	return buf.String(), nil
}

func postPaste(paste, api_key, user_key string) (string, error) {
	api_endpoint := "https://pastebin.com/api/api_post.php"

	request_dat := url.Values{}
	request_dat.Set("api_dev_key", api_key)
	request_dat.Set("api_paste_private", "2") // PRIVATE PASTE
	request_dat.Set("api_option", "paste")
	request_dat.Set("api_user_key", user_key)

	request_dat.Set("api_paste_code", paste)

	req, err := http.NewRequest("POST", api_endpoint, strings.NewReader(request_dat.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf := new(strings.Builder)
	io.Copy(buf, res.Body)

	if res.StatusCode != 200 {
		return "", errors.New("status code is not 200")
	}

	return buf.String(), nil
}
