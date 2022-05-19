package main

import (
	"fmt"
	"math"
	"os"
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
		panic(err)
	}
}

func die(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
