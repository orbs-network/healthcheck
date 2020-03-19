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
```

## Output JSON example

```json
{
  "status": "200 OK",
  "timestamp": "2020-03-19T11:50:21.0846185Z"
}
```

```json
{
  "error": "Get https://3123: dial tcp 0.0.12.51:443: connect: no route to host",
  "status": "Not Provisioned",
  "timestamp": "2020-03-19T13:52:34.842475+02:00"
}
```