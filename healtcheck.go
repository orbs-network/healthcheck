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

type Status struct {
	Status    string
	Timestamp time.Time
	Error     string
	Payload   map[string]interface{}
}

func defaultStatus(err error) Status {
	return Status{
		Status:    "Not provisioned",
		Timestamp: time.Now(),
		Error:     err.Error(),
	}
}

func Check(url string) (status Status, err error) {
	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		return defaultStatus(err), err
	}

	rawJSON, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return defaultStatus(err), err
	}

	jsonErr := json.Unmarshal(rawJSON, &status)
	if jsonErr != nil {
		wrappedErr := fmt.Errorf("failed to parse status JSON: %s", jsonErr)
		return defaultStatus(wrappedErr), wrappedErr
	}

	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return status, fmt.Errorf("unexpected response code %d", res.StatusCode)
	}

	return status, nil
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
	rawJSON, _ := json.MarshalIndent(status, "", "  ")

	os.MkdirAll(path.Dir(*output), 0644)
	ioutil.WriteFile(*output, rawJSON, 0644)

	if err != nil {
		fmt.Println(string(rawJSON))
		os.Exit(1)
	}
}
