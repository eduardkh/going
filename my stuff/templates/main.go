package main

import (
	"log"
	"os"
	"text/template"
)

/*
Learn and Use Templates in Go
https://levelup.gitconnected.com/learn-and-use-templates-in-go-aa6146b01a38
*/

func main() {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	type User struct {
		Name   string
		Coupon string
		Amount int64
	}
	user := User{
		Name:   "Rick",
		Coupon: "IAMAWESOMEGOPHER",
		Amount: 5000,
	}
	err = tpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
