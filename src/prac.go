package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	a := "Salam"
	b := "Shayan"
	c := fmt.Sprintf("%s, %s", a, b)
	fmt.Println(c)
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal("err: ", err)
	}

	fmt.Printf("\033[92m%s\n", resp.Status)
}
