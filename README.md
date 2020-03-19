# Node service health check

```sh
docker build -t orbsnetwork/healthcheck .
docker service create --name health --publish 8080:8080 orbsnetwork/healthcheck
```