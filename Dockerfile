FROM alpine:3.5

ADD . /go-ethereum
RUN \
  apk add --update git go make gcc musl-dev linux-headers && \
  (cd go-ethereum && make gwhale)                         && \
  cp go-ethereum/build/bin/gwhale /usr/local/bin/         && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /go-ethereum && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30373
EXPOSE 30373/udp

ENTRYPOINT ["gwhale"]
