package healthcheck

import (
	"context"
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
		err := fmt.Errorf("unexpected response code %d", res.StatusCode)
		status.Error = err.Error()
		return status, err
	}

	return status, nil
}

const WRITE_MODE = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
const APPEND_MODE = os.O_WRONLY | os.O_CREATE | os.O_APPEND

// os.O_WRONLY|os.O_CREATE|os.O_TRUNC for truncation
func WriteFile(filename string, data []byte, perm os.FileMode, openMode int) error {
	f, err := os.OpenFile(filename, openMode, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func DumpToDisk(ctx context.Context, output string, rawJSON []byte, mode int) {
	callback := make(chan interface{})
	defer close(callback)

	go func() {
		os.MkdirAll(path.Dir(output), 0644)
		if err := WriteFile(output, rawJSON, 0644, mode); err != nil {
			fmt.Println("failed to write to disk:", err.Error())
		}

		callback <- nil
	}()

	select {
	case <-ctx.Done():
		fmt.Println("failed to write to disk:", ctx.Err())
	case <-callback:
		return
	}
}

func Main() {
	url := flag.String("url", "", "url to query")
	output := flag.String("output", "", "path to file")
	log := flag.String("log", "", "path to log file")

	flag.Parse()

	if *url == "" || *output == "" {
		fmt.Println("url or output are missing")
		os.Exit(1)
	}

	status, err := Check(*url)
	rawJSON, _ := json.MarshalIndent(status, "", "  ")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// write to status
	DumpToDisk(ctx, *output, rawJSON, WRITE_MODE)

	if err != nil {
		// write to log in case of an error
		if *log != "" {
			rawJSON, _ := json.Marshal(status)
			rawJSON = append(rawJSON, []byte("\n")...)
			DumpToDisk(ctx, *log, rawJSON, APPEND_MODE)
		}

		fmt.Println(string(rawJSON))
		os.Exit(1)
	}
}
