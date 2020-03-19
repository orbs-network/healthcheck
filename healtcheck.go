package healthcheck

import (
	"fmt"
	"net/http"
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
