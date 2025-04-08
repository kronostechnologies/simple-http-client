package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: METHOD URL [BODY [CONTENT_TYPE]]")
		os.Exit(1)
	}

	method := os.Args[1]
	url := os.Args[2]
	var body string
	timeout := 5

	var contentType string
	if len(os.Args) >= 4 {
		body = os.Args[3]

		if len(os.Args) == 5 {
			contentType = os.Args[4]
		} else {
			contentType = "application/json"
		}
	}

	request, re := http.NewRequest(method, url, bytes.NewBufferString(body))
	if re != nil {
		log.Panicln(re)
	}

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	if et, lu := os.LookupEnv("HTTP_TIMEOUT"); lu {
		if it, ce := strconv.Atoi(et); ce == nil {
			timeout = it
		} else {
			log.Panicf("HTTP_TIMEOUT invalid: %s\n", et)
		}
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
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
