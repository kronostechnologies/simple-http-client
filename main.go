package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: METHOD URL [BODY]")
		os.Exit(1)
	}

	method := os.Args[1]
	url := os.Args[2]
	var body string

	if len(os.Args) == 4 {
		body = os.Args[3]
	}

	request, re := http.NewRequest(method, url, bytes.NewBufferString(body))
	if re != nil {
		log.Panicln(re)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, ce := client.Do(request)
	if ce != nil {
		log.Panicln(ce)
	}

	log.Printf("%s %s %d\n", method, url, response.StatusCode)

	if be := response.Body.Close(); be != nil {
		log.Panicln(be)
	}
}