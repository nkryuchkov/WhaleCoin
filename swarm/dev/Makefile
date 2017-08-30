.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/WhaleCoinOrg/WhaleCoin/cmd/swarm
	go build -o bin/geth     github.com/WhaleCoinOrg/WhaleCoin/cmd/geth
	go build -o bin/bootnode github.com/WhaleCoinOrg/WhaleCoin/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/WhaleCoinOrg/WhaleCoin/swarm/...
