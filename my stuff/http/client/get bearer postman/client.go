package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://httpbin.org/bearer"
	method := "GET"
	sendToken := "ZWRkaWV0aGVvbmUK"
	getToken, _ := getBearerToken(url, method, ``, sendToken)
	fmt.Println(getToken)
}
func getBearerToken(url string, method string, payload string, token string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", " application/json")
	req.Header.Add("Authorization", " Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
