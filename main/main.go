package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/orbs-network/healthcheck"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func main() {
	url := flag.String("url", "", "url to query")
	output := flag.String("output", "", "path to file")

	flag.Parse()

	if *url == "" || *output == "" {
		fmt.Println("url or output are missing")
		os.Exit(1)
	}

	status, err := healthcheck.Check(*url)

	result := map[string]interface{}{
		"timestamp": time.Now(),
		"status":    status,
	}

	if err != nil {
		result["error"] = err.Error()
	}

	rawJSON, _ := json.MarshalIndent(result, "", "  ")

	os.MkdirAll(path.Dir(*output), 0644)
	ioutil.WriteFile(*output, rawJSON, 0644)

	if err != nil {
		fmt.Println(string(rawJSON))
		os.Exit(1)
	}
}
