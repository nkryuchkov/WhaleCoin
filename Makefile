# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gwhale android ios gwhale-cross swarm evm all test clean
.PHONY: gwhale-linux gwhale-linux-386 gwhale-linux-amd64 gwhale-linux-mips64 gwhale-linux-mips64le
.PHONY: gwhale-linux-arm gwhale-linux-arm-5 gwhale-linux-arm-6 gwhale-linux-arm-7 gwhale-linux-arm64
.PHONY: gwhale-darwin gwhale-darwin-386 gwhale-darwin-amd64
.PHONY: gwhale-windows gwhale-windows-386 gwhale-windows-amd64

GOBIN = build/bin
GO ?= latest

gwhale:
	build/env.sh go run build/ci.go install ./cmd/gwhale
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gwhale\" to launch gwhale."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gwhale.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gwhale.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

gwhale-cross: gwhale-linux gwhale-darwin gwhale-windows gwhale-android gwhale-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-*

gwhale-linux: gwhale-linux-386 gwhale-linux-amd64 gwhale-linux-arm gwhale-linux-mips64 gwhale-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-*

gwhale-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gwhale
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep 386

gwhale-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gwhale
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep amd64

gwhale-linux-arm: gwhale-linux-arm-5 gwhale-linux-arm-6 gwhale-linux-arm-7 gwhale-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep arm

gwhale-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gwhale
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep arm-5

gwhale-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gwhale
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep arm-6

gwhale-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gwhale
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep arm-7

gwhale-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gwhale
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep arm64

gwhale-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gwhale
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep mips

gwhale-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gwhale
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep mipsle

gwhale-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gwhale
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep mips64

gwhale-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gwhale
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-linux-* | grep mips64le

gwhale-darwin: gwhale-darwin-386 gwhale-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-darwin-*

gwhale-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gwhale
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-darwin-* | grep 386

gwhale-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gwhale
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-darwin-* | grep amd64

gwhale-windows: gwhale-windows-386 gwhale-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-windows-*

gwhale-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gwhale
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-windows-* | grep 386

gwhale-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gwhale
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gwhale-windows-* | grep amd64
