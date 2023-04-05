package main

import (
	"11111/coincap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect")
			return nil
		},
	}
	for {
		resp, err := client.Get("https://api.coincap.io/v2/assets/")
		if err != nil {
			log.Fatal(nil)
		}
		defer resp.Body.Close()
		fmt.Println("Response status: ", resp.StatusCode)
		data1, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(nil)
		}
		cockies := resp.Cookies()
		fmt.Printf("cockies1: %v\n", cockies)
		var coincapApiStruct coincap.AssetsResponse
		if err := json.Unmarshal(data1, &coincapApiStruct); err != nil {
			log.Fatal(nil)
		}
		coincap.PrintInfo(&coincapApiStruct)
	}
}
