.PHONY: all build package osdsdock osdslet docker clean

all:build

build:osdsdock osdslet osdsctl

package:
	go get github.com/opensds/opensds/cmd/osdslet
	go get github.com/opensds/opensds/cmd/osdsdock
	go get github.com/opensds/opensds/osdsctl

osdsdock:package
	mkdir -p  ./build/out/bin/
	go build -o ./build/out/bin/osdsdock github.com/opensds/opensds/cmd/osdsdock

osdslet:package
	mkdir -p  ./build/out/bin/
	go build -o ./build/out/bin/osdslet github.com/opensds/opensds/cmd/osdslet

osdsctl:package
	mkdir -p  ./build/out/bin/
	go build -o ./build/out/bin/osdsctl github.com/opensds/opensds/osdsctl

docker:build
	cp ./build/out/bin/osdsdock ./cmd/osdsdock
	cp ./build/out/bin/osdslet ./cmd/osdslet
	docker build cmd/osdsdock -t opensdsio/opensds-dock:v1beta
	docker build cmd/osdslet -t opensdsio/opensds-controller:v1beta

clean:
	rm -rf ./build/out/bin/* ./cmd/osdslet/osdslet ./cmd/osdsdock/osdsdock
