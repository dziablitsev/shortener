package main

import (
	"flag"
)

func parseFlags() {
	flag.StringVar(&serverAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&shortURLHost, "b", "http://localhost:8080", "short url address")
	flag.BoolVar(&serverDebug, "debug", false, "debug mode")
	flag.IntVar(&shortURLLen, "len", 8, "short URL length")
	flag.Parse()
}
