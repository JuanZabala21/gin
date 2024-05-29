package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func test() {

	url := "localhost:8080/videos"
	method := "POST"

	payload := strings.NewReader(`{
    "title": "Practicando Gin Framework",
    "description": "Estudiando",
    "url": "www.youtube.com"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic anVhbjoxMjM0")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
