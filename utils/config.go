package utils

import (
	"fmt"

	"github.com/linkerd/linkerd2/testutil"
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
	return &ConformanceTestOptions{
		LinkerdVersion:     "stable-2.8.0",
		LinkerdNamespace:   "l5d-conformance",
		UpgradeFromVersion: "",
		ClusterDomain:      "cluster.local",
		ExternalIssuer:     false,
		Uninstall:          true,
		AddOns:             make(map[string]interface{}),
	}
}

func (options *ConformanceTestOptions) parseConfigValues() {
	var (
		defaultNs            = "l5d-conformance"
		defaultClusterDomain = "cluster.local"
		defaultVersion       = "stable-2.8.0"
	)
	if options.LinkerdVersion == "" {
		fmt.Println("Unspecified linkerd2 version - using default value \"stable-2.7.0\"")
		options.LinkerdVersion = defaultVersion

		// TODO: use the version API to fetch the latest stable instead of hard-coding values
	}

	if options.LinkerdNamespace == "" {
		fmt.Println("Unspecified linkerd2 control plane namespace - use default value \"l5d-conformance\"")
		options.LinkerdNamespace = defaultNs
	}

	if options.ClusterDomain == "" {
		fmt.Println("Unspecified cluster domain - using default value \"cluster.local\"")
		options.ClusterDomain = defaultClusterDomain
	}

}

func (options *ConformanceTestOptions) initNewTestHelperFromOptions() *testutil.TestHelper {
	var (
		linkerdPath = "/home/mayank/.linkerd2/bin/linkerd"

		//TODO: move these to ConformanceTestOptions while writing Helm tests
		helmPath        = "target/helm"
		helmChart       = "charts/linkerd2"
		helmStableChart = "linkerd/linkerd2"
		helmReleaseName = ""
	)

	return testutil.NewGenericTestHelper(
		linkerdPath,
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
