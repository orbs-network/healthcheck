package main

import "flag"

func main() {
	flag.String("url", "", "url to query")

	flag.String("output", "", "path to file")
}
