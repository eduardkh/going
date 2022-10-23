package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(100)
	log.Println(v)
}
