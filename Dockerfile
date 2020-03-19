FROM busybox

ADD _bin/* /opt/orbs/

WORKDIR /opt/orbs

HEALTHCHECK CMD /opt/orbs/healthcheck --url http://localhost:8080/metrics --output /opt/orbs/status/status.json

CMD /opt/orbs/service