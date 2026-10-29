package launchdarkly

import (
	"fmt"
	"path/filepath"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	ldprovider "github.com/launchdarkly/terraform-provider-launchdarkly/launchdarkly"
	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	"github.com/nellisauction/pulumi-launchdarkly/provider/pkg/version"
)

const (
	mainPkg = "launchdarkly"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-launchdarkly/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(ldprovider.NewPluginProvider(version.Version)()),
		Name:              "launchdarkly",
		Version:           version.Version,
		DisplayName:       "LaunchDarkly",
		Publisher:         "NellisAuction",
		PluginDownloadURL: "github://api.github.com/nellisauction/pulumi-launchdarkly",
		Description:       "A Pulumi package for managing LaunchDarkly resources.",
		Keywords:          []string{"pulumi", "launchdarkly", "category/utility"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/nellisauction/pulumi-launchdarkly",
		Repository:        "https://github.com/nellisauction/pulumi-launchdarkly",
		GitHubOrg:         "launchdarkly",
		Config: map[string]*tfbridge.SchemaInfo{
			"access_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"LAUNCHDARKLY_ACCESS_TOKEN"},
				},
				Secret: tfbridge.True(),
			},
			"oauth_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"LAUNCHDARKLY_OAUTH_TOKEN"},
				},
				Secret: tfbridge.True(),
			},
			"api_host": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"LAUNCHDARKLY_API_HOST"},
				},
			},
		},
		// Upstream does not mark these as sensitive, but they carry credentials
		// (e.g. Kinesis/mParticle keys, integration auth details).
		Resources: map[string]*tfbridge.ResourceInfo{
			"launchdarkly_audit_log_subscription": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"config": {Secret: tfbridge.True()},
				},
			},
			"launchdarkly_destination": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"config": {Secret: tfbridge.True()},
				},
			},
			"launchdarkly_integration_delivery_configuration": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"config": {Secret: tfbridge.True()},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@nellisauction/pulumi-launchdarkly",
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{RespectSchemaVersion: true}
			i.PyProject.Enabled = true
			return i
		})(),
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/nellisauction/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences:    map[string]string{"Pulumi": "3.*"},
			Namespaces:           map[string]string{mainPkg: "LaunchDarkly"},
		},
		MetadataInfo:                   tfbridge.NewProviderMetadata(metadata),
		EnableZeroDefaultSchemaVersion: true,
		EnableAccurateBridgePreview:    true,
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("launchdarkly_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
