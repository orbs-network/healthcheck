# Node service health check

Queries an endpoint and exits with code `1` unless it got a response with `200 OK`.

## Building

`./build-binaries.sh`

## CLI

```
Usage:
  -output string
    	path to file
  -url string
    	url to query
  -log string (optional)
        path to log file that keeps the statuses that contain errors
```

## Output JSON example

```json
{
  "Status": "200 OK",
  "Timestamp": "2020-03-19T11:50:21.0846185Z",
  "Payload": {
    "BlockStorage.BlockHeight": {
     "Name": "BlockStorage.BlockHeight",
     "Value": 3715964
    }
  }
}
```

```json
{
  "Error": "Get https://3123: dial tcp 0.0.12.51:443: connect: no route to host",
  "Status": "Not Provisioned",
  "Timestamp": "2020-03-19T13:52:34.842475+02:00"
}
```