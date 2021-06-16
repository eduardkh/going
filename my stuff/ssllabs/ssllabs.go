package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GOT IT HERE - https://stackoverflow.com/questions/43770273/json-unmarshalling-without-struct
func JsonIndent(jsontext []byte) ([]byte, error) {
	var err error
	var jsonIndent []byte
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(jsontext, &objmap)
	if err != nil {
		return jsonIndent, err
	}
	jsonIndent, err = json.MarshalIndent(objmap, "", "  ")
	return jsonIndent, err
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	url := "https://api.ssllabs.com/api/v3/analyze?host=digital.isracard.co.il&publish=on&latest=on"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	checkError(err)

	req.Header.Add("Cookie", "JSESSIONID=43032C0ABCC9F6AE967E8EBE241C572B")

	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)

	unmarshalByteBody, err := JsonIndent(body)
	checkError(err)
	fmt.Println(string(unmarshalByteBody))
}
