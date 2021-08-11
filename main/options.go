package main

import (
	flag "github.com/spf13/pflag"
)

var (
	// config
	configCmd = flag.NewFlagSet("foo", flag.ExitOnError)

	// start
	startCmd   = flag.NewFlagSet("bar", flag.ExitOnError)
	portOption = startCmd.Int("port", 0, "port")
)
