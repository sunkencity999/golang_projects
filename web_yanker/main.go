package main

import (
	"fmt"
	"log"
)
import "io/ioutil"
import "net/http"

func main() {
	res, err := http.Get("https://www.google.com/#q=Bay+Area+Realtor")
	page, err := ioutil.ReadAll(res.Body)
	filename := "aprhomepage.html"
	pageLog := ioutil.WriteFile(filename, []byte(page), 0666)

	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	fmt.Printf("%s", page)
	fmt.Printf("%s", pageLog)

}
