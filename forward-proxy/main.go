package main

import "github.com/Makepad-fr/sfp/core"

func main() {
	core.StartForwardProxy("localhost", "9090")
}
