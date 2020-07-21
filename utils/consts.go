package utils

const (
	defaultNs            = "l5d-conformance"
	defaultClusterDomain = "cluster.local"
	defaultPath          = "/.linkerd2/bin/linkerd"

	versionEndpointURL = "https://versioncheck.linkerd.io/version.json"

	//TODO: move these to ConformanceTestOptions while writing Helm tests
	helmPath        = "target/helm"
	helmChart       = "charts/linkerd2"
	helmStableChart = "linkerd/linkerd2"
	helmReleaseName = ""

	// TODO: move these to config while adding tests for multicluster
	multicluster                = false
	multiclusterHelmChart       = "multicluster-helm-chart"
	multiclusterHelmReleaseName = "multicluster-helm-release"

	installEnv           = "LINKERD2_VERSION"
	configFile           = "config.yaml"
	linkerdInstallScript = "install.sh"
	installScriptURL     = "https://run.linkerd.io/install"

	// string literals for identifying the ingress controllers

	// Nginx holds the string literal "nginx"
	Nginx = "nginx"

	// Traefik    = "traefik"
	// GCE        = "gce"
	// Ambassador = "ambassador"
	// Gloo       = "gloo"
	// Contour    = "contour"

	// NginxNs is the namespace in which the nginx controller is installed
	NginxNs = "ingress-nginx"

	// NginxController is the name of the nginx controller
	NginxController = "ingress-nginx-controller"
)
