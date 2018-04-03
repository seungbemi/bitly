package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	flag.Parse()
}

type data struct {
	URL     string `json:"url"`
	Hash    string `json:"hash"`
	GHash   string `json:"global_hash"`
	LongURL string `json:"long_url"`
}

type respBody struct {
	Contents data `json:"data"`
}

type respStatus struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status_txt"`
}

func main() {
	longURL := flag.Arg(0)
	apikey := os.Getenv("API_KEY")
	user := os.Getenv("USER")

	if len(longURL) == 0 {
		fmt.Println("Error : URL is empty")
		return
	}
	if len(apikey) == 0 {
		fmt.Println("Error : apikey is set")
		return
	}
	if len(user) == 0 {
		fmt.Println("Error : user is set")
		return
	}

	resp, err := http.Get(fmt.Sprintf("https://api-ssl.bitly.com/v3/shorten?format=json&login=%s&apiKey=%s&longUrl=%s", user, apikey, longURL))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error : " + err.Error())
		return
	}

	status := respStatus{}
	if err := json.Unmarshal(body, &status); err != nil {
		fmt.Println("Error : " + err.Error())
		return
	}

	if status.StatusCode != http.StatusOK {
		fmt.Println("Error : " + status.Status)
		return
	}

	url := respBody{}
	if err := json.Unmarshal(body, &url); err != nil {
		fmt.Println("Error : " + err.Error())
		return
	}

	fmt.Println(url.Contents.URL)
}
