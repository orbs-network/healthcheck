package healthcheck

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

func Check(url string) (status string, err error) {
	status = "Not Provisioned"

	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	status = res.Status
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		err = fmt.Errorf("wrong response")
	}

	return
}

func Main() {
	url := flag.String("url", "", "url to query")
	output := flag.String("output", "", "path to file")

	flag.Parse()

	if *url == "" || *output == "" {
		fmt.Println("url or output are missing")
		os.Exit(1)
	}

	status, err := Check(*url)

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
