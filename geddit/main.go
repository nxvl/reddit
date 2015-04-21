package main

import (
	"fmt"
	"github.com/nxvl/reddit"
	"log"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
