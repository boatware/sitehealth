package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var format = ""
var nl = "\n"
var loop = false

// 1000 ms as default
var delay int64 = 1000 * 1000000

func microTime() float64 {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	micSeconds := float64(now.Nanosecond()) / 1000000000
	return float64(now.Unix()) + micSeconds
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func die(msg string) {
	log.Println(msg)
	os.Exit(1)
}

func detectOS() {
	if //goland:noinspection ALL
	runtime.GOOS == "windows" {
		nl = "\r\n"
	}
}

func processArgs() {
	if len(os.Args) < 2 {
		die("Not enough arguments.")
	}

	if len(os.Args) > 2 {
		for _, arg := range os.Args[2:] {
			if arg == "--json" {
				format = "json"
			}

			if arg == "--csv" {
				format = "csv"
			}

			if arg == "--loop" {
				loop = true
			}

			delayMatch, _ := regexp.Match(`^--delay=\d.*$`, []byte(arg))
			if delayMatch {
				d := strings.ReplaceAll(arg, "--delay=", "")
				dInt, _ := strconv.ParseInt(d, 10, 64)
				delay = dInt * 1000000
			}
		}
	}
}

func getURL() string {
	url := os.Args[1]
	match, _ := regexp.Match(`^http(s|)://\w.*\.\w{2,3}$`, []byte(url))
	if !match {
		die("Argument '" + url + "' is not a valid URL")
	}

	return url
}

func pingURL(url string) {
	start := microTime()
	resp, err := http.Get(url)
	check(err)

	res := (microTime() - start) * 1000

	template := "%s\t%3d\t%f\n"
	switch format {
	case "json":
		template = "{\"url\":\"%s\",\"status\":%3d,\"time\":%f}"
		break

	case "csv":
		if !loop {
			template = "url,status,time" + nl + "%s,%3d,%f" + nl
		} else {
			template = "%s,%3d,%f" + nl
		}
	}

	fmt.Printf(template, url, resp.StatusCode, round(res, 000000.000001))
}

func main() {
	detectOS()
	processArgs()
	url := getURL()

	if !loop {
		pingURL(url)
		os.Exit(0)
	}

	for {
		pingURL(url)
		time.Sleep(time.Duration(delay))
	}
}
