package main

import (
	provider "github.com/nellisauction/pulumi-launchdarkly/provider"
	pftfgen "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"
)

func main() {
	pftfgen.Main("launchdarkly", provider.Provider())
}
