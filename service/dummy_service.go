package service

import "net/http"

func Dummy() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
  "Timestamp": "2020-06-30T15:48:18.434309+03:00",
  "Status": "Last Successful Committed Block was too long ago",
  "Error": "",
  "Payload": {
    "Uptime": 25,
    "BlockStorage_BlockHeight": 501,
    "StateStorage_BlockHeight": 501,
    "BlockStorage_LastCommitted": 1574694046590916000,
    "Gossip_IncomingConnections": 0,
    "Gossip_OutgoingConnections": 0,
    "Management_LastUpdated": 0,
    "Management_Subscription": "Active",
    "Version": {
      "Semantic": "",
      "Commit": ""
    }
  }
}`))
	})

	http.HandleFunc("/status.500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{
  "Timestamp": "2020-06-30T15:48:18.434309+03:00",
  "Status": "Last Successful Committed Block was too long ago",
  "Error": "",
  "Payload": {
    "Uptime": 25,
    "BlockStorage_BlockHeight": 501,
    "StateStorage_BlockHeight": 501,
    "BlockStorage_LastCommitted": 1574694046590916000,
    "Gossip_IncomingConnections": 0,
    "Gossip_OutgoingConnections": 0,
    "Management_LastUpdated": 0,
    "Management_Subscription": "Active",
    "Version": {
      "Semantic": "",
      "Commit": ""
    }
  }
}
`))
	})

	http.HandleFunc("/status.failed", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte{1, 2, 3})
	})
}
