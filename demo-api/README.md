# Demo-Api Service

This is the Demo-Api service

Generated with

```
micro new github.com/go-lego/demo-api --namespace=go.lego --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.lego.api.demo-api
- Type: api
- Alias: demo-api

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./demo-api-api
```

Build a docker image
```
make docker
```