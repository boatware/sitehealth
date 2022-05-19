package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processArgs() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	if len(os.Args) >= 1 {
		for _, arg := range os.Args[1:] {
			if arg == "--help" || arg == "-h" {
				usage()
				os.Exit(0)
			}

			if arg == "--json" {
				format = "json"
			}

			if arg == "--csv" {
				format = "csv"
			}

			if arg == "--loop" {
				loop = true
				if followRedirect {
					die("Cannot follow redirects in a loop")
				}
			}

			if arg == "--include-length" || arg == "-l" {
				getContentLength = true
			}

			if arg == "--follow-redirect" || arg == "-f" {
				followRedirect = true
				if loop {
					die("Cannot follow redirects in a loop")
				}
			}

			delayMatch, _ := regexp.Match(`^--delay=\d.*$`, []byte(arg))
			if delayMatch {
				d := strings.ReplaceAll(arg, "--delay=", "")
				dInt, _ := strconv.ParseInt(d, 10, 64)
				delay = dInt * 1000000
			}

			urlMatch, _ := regexp.Match(`^http(s|)://\w.*\.\w{2,3}$`, []byte(arg))
			if urlMatch {
				url = arg
			}
		}

		if url == "" {
			die("No URL given.")
		}
	}
}
