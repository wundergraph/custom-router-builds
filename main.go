package main

import (
	routercmd "github.com/wundergraph/cosmo/router/cmd"
	// Import your modules here using "_ <package> syntax". Example -->
	_ "github.com/wundergraph/custom-router-builds/moduletemplate"
)

func main() {
	routercmd.Main()
}
