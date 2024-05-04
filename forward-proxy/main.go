package main

import (
	"log"
	"os"

	"github.com/Makepad-fr/sfp/core"
)

var hostname = ""
var port = "9090"
var config core.Config

func init() {
	val, ok := os.LookupEnv("SFP_HOSTNAME")
	if ok {
		hostname = val
	}
	val, ok = os.LookupEnv("SFP_PORT")
	if ok {
		port = val
	}
	var err error
	if len(os.Args) > 1 {
		config, err = core.LoadConfigFromFile(os.Args[1])
		if err != nil {
			log.Fatalf("Error while loading configuration file %s: %v", os.Args[1], err)
		}
		return
	}
	val, ok = os.LookupEnv("SFP_CONFIG_FILE")
	if ok {
		config, err = core.LoadConfigFromFile(val)
		if err != nil {
			log.Fatalf("Error while loading configuration file %s: %v", val, err)
		}
		return
	}
	config = core.Config{
		Tls: nil,
		ServerAddress: core.ServerAddress{
			HostName: hostname,
			Port:     port,
		},
		AccessLogging: nil,
	}

}

func main() {
	core.Start(config)
}
