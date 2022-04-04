package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type JSTModel struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var jstData JSTModel

func main() {
	jsonData, err := os.ReadFile("json.json")
	check(err)
	err = json.Unmarshal(jsonData, &jstData)
	check(err)
	fmt.Printf("ID %d\nUser ID %d\nTitle %s\ncompleted %v", jstData.ID, jstData.UserID, jstData.Title, jstData.Completed)
}
