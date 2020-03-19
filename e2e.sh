#!/bin/bash -x

./build-binaries.sh
docker build -t orbsnetwork/healthcheck .

rm -rf tmp
mkdir -p tmp

docker service rm health
docker service create --name health --publish 8080:8080 --mount type=bind,source=$(pwd)/tmp,destination=/opt/orbs/status orbsnetwork/healthcheck

while true; do
    if [ -f ./tmp/status.json ]; then
        cat ./tmp/status.json
        break
    else
        sleep 1
    fi
done
