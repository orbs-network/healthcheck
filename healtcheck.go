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
	Error     error
	Payload   map[string]interface{}
}

func Check(url string) (status Status, err error) {
	status.Status = "Not Provisioned"
	status.Timestamp = time.Now()

	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		status.Error = err
		return
	}
	rawJSON, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		status.Error = err
		return
	}

	status.Status = res.Status
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		err = fmt.Errorf("unexpected response code %d", res.StatusCode)
		status.Error = err
	}

	payload := make(map[string]interface{})
	jsonErr := json.Unmarshal(rawJSON, &payload)
	if jsonErr != nil {
		println(err)
	}

	status.Payload = payload

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
	rawJSON, _ := json.MarshalIndent(status, "", "  ")

	os.MkdirAll(path.Dir(*output), 0644)
	ioutil.WriteFile(*output, rawJSON, 0644)

	if err != nil {
		fmt.Println(string(rawJSON))
		os.Exit(1)
	}
}
