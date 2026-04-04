package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"

	stripe "github.com/humanlayer/pulumi-stripe/provider"
	"github.com/humanlayer/pulumi-stripe/provider/pkg/version"
)

func main() {
	tfbridge.Main("stripe", version.Version, stripe.Provider(), nil)
}
