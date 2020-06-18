package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/linkerd/linkerd2/testutil"
)

var (
	defaultNs            = "l5d-conformance"
	defaultClusterDomain = "cluster.local"
	defaultVersion       = "stable-2.8.0"
	defaultPath          = "/.linkerd2/bin/linkerd"
)

// ConformanceTestOptions holds the values fed from the test config file
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

func getDefaultConformanceOptions() (*ConformanceTestOptions, error) {
	options := ConformanceTestOptions{}
	if err := options.parseConfigValues(); err != nil {
		return nil, err
	}
	options.Uninstall = true // uninstall by defaut

	return &options, nil
}

func getDefaultLinkerdPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/.linkerd2/bin/linkerd", home), nil
}

func (options *ConformanceTestOptions) parseConfigValues() error {
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
		path, err := getDefaultLinkerdPath()
		if err != nil {
			return err
		}
		fmt.Printf("Unspecified path to linkerd2 binary - using default value \"%s\"\n", path)
		options.LinkerdBinaryPath = path
	}

	return nil
}

func initK8sHelper(context string, retryFor func(time.Duration, func() error) error) (*testutil.KubernetesHelper, error) {
	k8sHelper, err := testutil.NewKubernetesHelper(context, retryFor)
	if err != nil {
		return nil, err
	}

	return k8sHelper, nil
}

func (options *ConformanceTestOptions) initNewTestHelperFromOptions() (*testutil.TestHelper, error) {
	var (
		//TODO: move these to ConformanceTestOptions while writing Helm tests
		helmPath        = "target/helm"
		helmChart       = "charts/linkerd2"
		helmStableChart = "linkerd/linkerd2"
		helmReleaseName = ""
	)

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	helper := testutil.NewGenericTestHelper(
		options.LinkerdBinaryPath,
		options.LinkerdVersion,
		options.LinkerdNamespace,
		options.UpgradeFromVersion,
		options.ClusterDomain,
		helmPath,
		helmChart,
		helmStableChart,
		helmReleaseName,
		options.ExternalIssuer,
		options.Uninstall,
		httpClient,
		testutil.KubernetesHelper{},
	)

	k8sHelper, err := initK8sHelper(options.K8sContext, helper.RetryFor)
	if err != nil {
		return nil, fmt.Errorf("error initializing k8s helper: %s", err)
	}

	helper.KubernetesHelper = *k8sHelper
	return helper, nil
}
