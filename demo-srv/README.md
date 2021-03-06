# Demo-Srv Service

This is the Demo-Srv service

Generated with

```
micro new github.com/go-lego/demo-srv --namespace=go.lego --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.lego.srv.demo-srv
- Type: srv
- Alias: demo-srv

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
./demo-srv-srv
```

Build a docker image
```
make docker
```