package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"time"
)

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

func main() {
	start := microTime()
	if len(os.Args) < 2 {
		die("Not enough arguments.")
	}

	url := os.Args[1]
	match, _ := regexp.Match(`^http(s|)://\w.*`, []byte(url))
	if match {
		resp, err := http.Get(url)
		check(err)

		res := (microTime() - start) * 1000
		fmt.Printf("%s\t%3d\t%f\n", url, resp.StatusCode, round(res, 000000.000001))
	} else {
		die("Argument '" + url + "' is not a valid URL")
	}
}
