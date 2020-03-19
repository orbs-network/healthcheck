package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
  "BlockStorage.BlockHeight": {
    "Name": "BlockStorage.BlockHeight",
    "Value": 3715964
  }
}`))
	})

	http.ListenAndServe(":8080", nil)
}