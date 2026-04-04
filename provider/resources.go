package stripe

import (
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	stripeshim "github.com/stripe/terraform-provider-stripe"

	"github.com/humanlayer/pulumi-stripe/provider/pkg/version"
)

//go:embed cmd/pulumi-resource-stripe/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 shimv2.NewProvider(stripeshim.NewProvider()),
		Name:              "stripe",
		DisplayName:       "Stripe",
		Publisher:         "HumanLayer",
		Version:           version.Version,
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		PluginDownloadURL: "github://api.github.com/humanlayer/pulumi-stripe",
		Description:       "A Pulumi provider for managing Stripe resources",
		Keywords:          []string{"pulumi", "stripe", "payments", "billing"},
		License:           "MIT",
		Homepage:          "https://github.com/humanlayer/pulumi-stripe",
		Repository:        "https://github.com/humanlayer/pulumi-stripe",
		GitHubOrg:         "humanlayer",
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources:         map[string]*tfbridge.ResourceInfo{},
		DataSources:       map[string]*tfbridge.DataSourceInfo{},
	}

	prov.MustComputeTokens(tokens.SingleModule("stripe_", "index", tokens.MakeStandard("stripe")))
	prov.SetAutonaming(255, "-")

	return prov
}
