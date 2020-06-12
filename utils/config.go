package utils

import (
	"fmt"

	"github.com/linkerd/linkerd2/testutil"
)

var (
	defaultNs            = "l5d-conformance"
	defaultClusterDomain = "cluster.local"
	defaultVersion       = "stable-2.8.0"
	defaultPath          = "/.linkerd2/bin/linkerd"
)

type ConformanceTestOptions struct {
	LinkerdVersion     string                 `yaml:"linkerdVersion,omitempty"`
	LinkerdNamespace   string                 `yaml:"linkerdNamespace,omitempty"`
	LinkerdBinaryPath  string                 `yaml:"linkerdPath,omitempty"`
	UpgradeFromVersion string                 `yaml:"upgradeFromVersion,omitempty"`
	ClusterDomain      string                 `yaml:"clusterDomain,omitempty"`
	K8sContext         string                 `yaml:"k8sContext,omitempty"`
	ExternalIssuer     bool                   `yaml:"externalIssuer,omitempty"`
	Uninstall          bool                   `yaml:"uninstall,omitempty"`
	AddOns             map[string]interface{} `yaml:"addOns,omitempty"`

	// TODO: Add fields for test specific configurations
	// TODO: Add fields for Helm tests
}

func getDefaultConformanceOptions() *ConformanceTestOptions {
	options := ConformanceTestOptions{}
	options.parseConfigValues()
	options.Uninstall = true // uninstall by defaut

	return &options
}

func (options *ConformanceTestOptions) parseConfigValues() {
	if options.LinkerdVersion == "" {
		fmt.Printf("Unspecified linkerd2 version - using default value \"%s\"\n", defaultVersion)
		options.LinkerdVersion = defaultVersion

		// TODO: use the version API to fetch the latest stable instead of hard-coding values
	}

	if options.LinkerdNamespace == "" {
		fmt.Printf("Unspecified linkerd2 control plane namespace - use default value \"%s\"\n", defaultNs)
		options.LinkerdNamespace = defaultNs
	}

	if options.ClusterDomain == "" {
		fmt.Printf("Unspecified cluster domain - using default value \"%s\"\n", defaultClusterDomain)
		options.ClusterDomain = defaultClusterDomain
	}

	if options.LinkerdBinaryPath == "" {
		fmt.Printf("Unspecified path to linkerd2 binary - using default value \"%s\"\n", defaultPath)
	}

}

func (options *ConformanceTestOptions) initNewTestHelperFromOptions() *testutil.TestHelper {
	var (
		//TODO: move these to ConformanceTestOptions while writing Helm tests
		helmPath        = "target/helm"
		helmChart       = "charts/linkerd2"
		helmStableChart = "linkerd/linkerd2"
		helmReleaseName = ""
	)

	return testutil.NewGenericTestHelper(
		options.LinkerdBinaryPath,
		options.LinkerdNamespace,
		options.UpgradeFromVersion,
		options.ClusterDomain,
		helmPath,
		helmChart,
		helmStableChart,
		helmReleaseName,
		options.ExternalIssuer,
		options.Uninstall,
	)
}
