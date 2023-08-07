package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/ksoclabs/terraform-provider-ksoc/internal/ksoc"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	Version string = "dev"
	Commit  string = ""
)

func main() {
	debugMode := flag.Bool("debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        *debugMode,
		ProviderAddr: "registry.terraform.io/ksoclabs/ksoc",
		ProviderFunc: ksoc.New(Version),
	}

	plugin.Serve(opts)
}
