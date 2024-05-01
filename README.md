# Simple Forward Proxy

This is a very simplistic implementation of forward proxy in Go programming language.

## Prerequisites

- Go 1.22 or higher
- Make

## Usage

### Build the project

```bash
make
```

### Start the forward proxy server

```bash
./out/forward-proxy
```

### Use the proxy server in a CURL

```bash
 curl -x http://localhost:9090 https://google.com
```