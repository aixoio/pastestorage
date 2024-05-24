package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/aixoio/pastestorage/aes"
	"github.com/aixoio/pastestorage/converter"
	"github.com/aixoio/pastestorage/hashing"
)

func DownloadFile(api_key, username, password, aes_key, link string) {

	user_key, _ := loginPost(username, password, api_key)
	p, _ := getPaste(strings.ReplaceAll(link, "https://pastebin.com/", ""), api_key, user_key)

	key := hashing.Sha256_to_bytes([]byte(aes_key))

	dat, _ := base64.RawStdEncoding.DecodeString(p)
	dec, _ := aes.AesGCMDecrypt(key, dat)

	var info final_dat

	json.Unmarshal(dec, &info)

	sort.Slice(info.Links, func(i, j int) bool {
		return info.Links[i].Index < info.Links[j].Index
	})

	all_dat_Str_s := []string{}

	for _, link := range info.Links {
		dat, _ := getPaste(strings.ReplaceAll(link.Link, "https://pastebin.com/", ""), api_key, user_key)
		all_dat_Str_s = append(all_dat_Str_s, dat)
	}

	converter.ConvertTextToFile(info.Filename, aes_key, all_dat_Str_s)

}

func getPaste(link, api_key, user_key string) (string, error) {
	api_endpoint := "https://pastebin.com/api/api_post.php"

	request_dat := url.Values{}
	request_dat.Set("api_dev_key", api_key)
	request_dat.Set("api_option", "show_paste")
	request_dat.Set("api_user_key", user_key)
	request_dat.Set("api_paste_key", link)

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
		fmt.Println(res.StatusCode)
		return "", errors.New("status code is not 200")
	}

	return buf.String(), nil
}
