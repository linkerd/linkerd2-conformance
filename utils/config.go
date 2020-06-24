package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	versionEndpointURL = "https://versioncheck.linkerd.io/version.json"
)

// Inject holds the inject test configuration
type Inject struct {
	Skip  bool `yaml:"skip,omitempty"`
	Clean bool `yaml:"clean,omitempty"` // deletes all resources created while testing
}

// GlobalControlPlane holds the options for installing a single control plane
type GlobalControlPlane struct {
	Enable    bool `yaml:"enable,omitempty"`
	Uninstall bool `yaml:"uninstall,omitempty"`
}

// ConformanceTestOptions holds the values fed from the test config file
type ConformanceTestOptions struct {
	LinkerdVersion     string                 `yaml:"linkerdVersion,omitempty"`
	LinkerdNamespace   string                 `yaml:"linkerdNamespace,omitempty"`
	LinkerdBinaryPath  string                 `yaml:"linkerdPath,omitempty"`
	UpgradeFromVersion string                 `yaml:"upgradeFromVersion,omitempty"`
	ClusterDomain      string                 `yaml:"clusterDomain,omitempty"`
	K8sContext         string                 `yaml:"k8sContext,omitempty"`
	ExternalIssuer     bool                   `yaml:"externalIssuer,omitempty"`
	AddOns             map[string]interface{} `yaml:"addOns,omitempty"`
	GlobalControlPlane `yaml:"globalControlPlane"`
	Inject             `yaml:"inject"`

	// TODO: Add fields for test specific configurations
	// TODO: Add fields for Helm tests
}

func getLatestStableVersion() (string, error) {

	var versionResp map[string]string

	req, err := http.NewRequest("GET", versionEndpointURL, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(body, &versionResp); err != nil {
		return "", err
	}

	return versionResp["stable"], nil
}

func getDefaultLinkerdPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/.linkerd2/bin/linkerd", home), nil
}

func initK8sHelper(context string, retryFor func(time.Duration, func() error) error) (*testutil.KubernetesHelper, error) {
	k8sHelper, err := testutil.NewKubernetesHelper(context, retryFor)
	if err != nil {
		return nil, err
	}

	return k8sHelper, nil
}

func (options *ConformanceTestOptions) parseDefaultConfigValues() error {
	if options.LinkerdVersion == "" {
		var version string
		var err error

		if version, err = getLatestStableVersion(); err != nil {
			fmt.Printf("error fetching latest version: %s\n", err)
			version = defaultVersion
		}

		fmt.Printf("Unspecified linkerd2 version - using default value \"%s\"\n", version)
		options.LinkerdVersion = version
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

	if !options.GlobalControlPlane.Enable && options.GlobalControlPlane.Uninstall {
		fmt.Println("globalControlPlane.uninstall will be ignored as globalControlPlane is disabled")
		options.GlobalControlPlane.Uninstall = false
	}

	return nil
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
		options.GlobalControlPlane.Uninstall,
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
