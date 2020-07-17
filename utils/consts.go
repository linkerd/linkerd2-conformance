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
	Nginx      = "nginx"
	Traefik    = "traefik"
	GCE        = "gce"
	Ambassador = "ambassador"
	Gloo       = "gloo"
	Contour    = "contour"

	NginxNs         = "ingress-nginx"
	NginxController = "ingress-nginx-controller"
)
