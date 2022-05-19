package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sitehealth/template"
	"time"
)

var format = ""
var nl = "\n"
var loop = false
var version = "1.0.0"
var url = ""
var getContentLength = false
var followRedirect = false

// 1000 ms as default
var delay int64 = 1000 * 1000000

func detectOS() {
	if //goland:noinspection ALL
	runtime.GOOS == "windows" {
		nl = "\r\n"
	}
}

func pingURL() {
	start := microTime()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(url)
	check(err)

	statusCode := resp.StatusCode
	responseTime := (microTime() - start) * 1000
	responseTimeRounded := round(responseTime, 0.000001)

	var destination string
	location, err := resp.Location()
	if err != nil {
		destination = url
	} else {
		destination = location.String()
	}

	var contentLength int64 = -1
	if getContentLength {
		contentLength = resp.ContentLength
	}

	t := template.Render(format, url, statusCode, responseTimeRounded, destination, contentLength)
	fmt.Print(t + nl)

	if followRedirect && !loop && statusCode < 400 && statusCode >= 300 {
		url = destination
		pingURL()
	}
}

func main() {
	detectOS()
	processArgs()

	if !loop {
		pingURL()
		os.Exit(0)
	}

	for {
		pingURL()
		time.Sleep(time.Duration(delay))
	}
}
