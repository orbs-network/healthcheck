package service

import "net/http"

func Dummy() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
  "BlockStorage.BlockHeight": {
    "Name": "BlockStorage.BlockHeight",
    "Value": 3715964
  }
}`))
	})

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
}
