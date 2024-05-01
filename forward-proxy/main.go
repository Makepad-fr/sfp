package main

import (
	"os"

	"github.com/Makepad-fr/sfp/core"
)

var hostname = ""
var port = "9090"

func init() {
	val, ok := os.LookupEnv("SFP_HOSTNAME")
	if ok {
		hostname = val
	}
	val, ok = os.LookupEnv("SFP_PORT")
	if ok {
		port = val
	}
}

func main() {
	core.StartForwardProxy(hostname, port)
}
