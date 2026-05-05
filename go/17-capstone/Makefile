.PHONY: build test e2e lint clean

BINARY := kubetail
PKG    := example.com/golab/17-capstone

VERSION ?= dev
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE    ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS := -X '$(PKG)/internal/version.Version=$(VERSION)' \
           -X '$(PKG)/internal/version.Commit=$(COMMIT)' \
           -X '$(PKG)/internal/version.Date=$(DATE)'

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/kubetail

test:
	go test -race ./...

e2e:
	@echo "Requires kind + kubectl in PATH. See README Stage walkthrough."
	@command -v kind >/dev/null || (echo "kind not installed: https://kind.sigs.k8s.io/" && exit 1)
	@command -v kubectl >/dev/null || (echo "kubectl not installed" && exit 1)
	kind get clusters 2>/dev/null | grep -q kubetail-lab || kind create cluster --name kubetail-lab
	kubectl apply -f testdata/loggy.yaml
	kubectl rollout status deploy/loggy --timeout=60s
	@echo "Streaming for 10s..."
	@timeout 10s ./$(BINARY) -n default -l app=loggy --no-color || true

lint:
	gofmt -l .
	go vet ./...
	@command -v golangci-lint >/dev/null && golangci-lint run ./... || echo "golangci-lint not installed; skipping"

clean:
	rm -f $(BINARY)
