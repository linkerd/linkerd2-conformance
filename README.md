![Linkerd][logo]

This repo contains the conformance tests for [Linkerd2](https://github.com/linkerd/linkerd2) as described by [this](https://github.com/linkerd/rfc/pull/24) RFC.

The conformance validation tool is primarily intended to be run on a specified version of Linkerd to verify the correctness of a [Kubernetes](https://kubernetes.io/) cluster's configuration with respect to Linkerd as well as validate non-trivial network communication (HTTP, gRPC, websocket) among stateless and stateful workloads in the Linkerd data plane.

The conformance tests exercise the following features:

- [ ] Validation of your Linkerd2 control plane
- [ ] Automatic proxy injection on workloads
- [ ] Functioning of  `linkerd tap`, `stat`, `routes` and `edges` commands
- [ ] Verifying the functioning of the `tap` extension API server
- [ ] Retries and timeouts
- [ ] Data plane health checks
- [ ] Ingress configuration

_...and much more_

If you are interested in helping to extend the test suites, see [Adding new tests](https://github.com/linkerd/linkerd2-conformance#2-adding-tests) below.


## Table of Contents

- [Repository Structure](https://github.com/mayankshah1607/linkerd2-conformance#repository-structure)
- [Configuring your tests](https://github.com/mayankshah1607/linkerd2-conformance#configuring-your-tests)
  - [Configuration Options](https://github.com/mayankshah1607/linkerd2-conformance#configuration-options)
- [Usage](https://github.com/mayankshah1607/linkerd2-conformance#usage)
  - [Using the Sonobuoy CLI](https://github.com/mayankshah1607/linkerd2-conformance#using-the-sonobuoy-cli)
  - [Running the tests using Docker]()
  - [Running the tests locally](https://github.com/mayankshah1607/linkerd2-conformance#running-the-tests-locally)
- [Adding new tests](https://github.com/mayankshah1607/linkerd2-conformance#adding-new-tests)
  - [Initial bootstrapping](https://github.com/mayankshah1607/linkerd2-conformance#1-initial-bootstrapping)
  - [Adding tests](https://github.com/mayankshah1607/linkerd2-conformance#2-adding-tests)
  - [Wiring up newly added tests](https://github.com/mayankshah1607/linkerd2-conformance#3-wiring-up-newly-added-test)

## Repository Structure

- [`tests`](https://github.com/mayankshah1607/linkerd2-conformance/tree/master/specs) contains the tests for each of the features organized into separate packages
- [`sonobuoy`](https://github.com/mayankshah1607/linkerd2-conformance/tree/master/sonobuoy) contains the items required to be able to run the tests as a [Sonobuoy](https://github.com/vmware-tanzu/sonobuoy) plugin
- [`utils`](https://github.com/mayankshah1607/linkerd2-conformance/tree/master/utils) contains helper functions that can be used while writing conformance tests
- [`bin`](https://github.com/mayankshah1607/linkerd2-conformance/blob/master/bin) contains useful helper scripts to build/push the Docker image and running the tests
- [`test_template`](https://github.com/mayankshah1607/linkerd2-conformance/tree/master/test_template) contains a sample template for adding new tests under the `tests/` folder

## Configuring your tests

The conformance tests can be easily configured by specifying a `config.yaml` that holds the configuration values. This includes things like Linkerd version, Linkerd add-ons, test specific configuration, binary path, etc.  

Run the command below to pull up a sample configuration file, `config.yaml`, and modify it according to your requirements. The tests shall read this YAML file during runtime and run accordingly.

```bash
$ curl -sL https://raw.githubusercontent.com/mayankshah1607/linkerd2-conformance/master/config.yaml > config.yaml
```

### Configuration options

This section describes the various configuration options and its default values. Not providing a configuration file, or providing a partial configuration file will result in the tests running on default settings as described below

| Option | Description | Default value |
|-|-|-|
| `linkerdVersion` | The linkerd2 binary version to use | Latest stable release |
| `linkerdBinaryPath` | If specified, the tests use the binary installed in the directory. It is recommended that this is left unspecified while using Sonobuoy or if upgrade tests are enabled | `$HOME/.linkerd2/bin/linkerd` | 
| `clusterDomain` | Use the specified cluster domain | `"cluster.local"` |
| `K8sContext` | Use the specified K8s context. Its is recommended that while running the tests with Sonobuoy (`sonobuoy run`), use the `--context` flag | `""` |
| `controlPlane.namespace` | Installs the control plane in the specified namespace | `"l5d-conformance"` |
| `controlPlane.config.ha` | Use a high-availability control plane for the tests | `false` |
| `controlPlane.config.flags` | Use the specified `linkerd install` CLI flag options while testing control plane installation | `[]` |
| `controlPlane.config.addOns` | Use the specified add-on configuration while testing control plane installation | `nil` |
| `testCase.lifecycle.skip` | Skip the pre-flight control plane installation tests | `false` |
| `testCase.lifecycle.upgradeFromVersion` | If specified, first install the CLI and control plane using the specified version, and test if they can be upgraded to `linkerdVersion` | `""` |
| `testCase.lifecycle.reinstall` | If true, install a new control plane for each test. Otherwise, use a single control plane throughout | `false` |
| `testCase.lifecycle.uninstall` | If using a single control plane, uninstall once the tests complete (whether they pass or fail) | `false` |
| `testCase.inject.skip` | Skip proxy injection tests | `false` |
| `testCase.inject.clean` | Delete the resources created for testing proxy injection | `false` |
| `testCase.ingress.skip` | If true, skips all ingress tests | `false` |
| `testCase.ingress.config.controllers` | List of ingress controllers to test. Currently only supports `nginx` | []string |


## Usage

This section outlines the various methods that can be used to run the conformance tests against your [Kubernetes](https://kubernetes.io/) cluster

### Using the Sonobuoy CLI

[Sonobuoy](https://github.com/vmware-tanzu/sonobuoy) offers a reliable way to run diagnostic tests in a Kubernetes cluster. We leverage its [plugin model](https://sonobuoy.io/docs/master/plugins/) to run conformance tests inside a Kubernetes pod.


The below commands assume that the user has the [Sonobuoy CLI](https://github.com/vmware-tanzu/sonobuoy#installation) and [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) (with a correctly configured `kubeconfig`) installed locally.

This repo provides a Sonobuoy plugin file that is intended to be plugged into the Sonobuoy CLI. Sonobuoy reads the plugin definition, and spins up a pod with the [linkerd2-conformance Docker image]().

```bash
# Create a ConfigMap from the `config.yaml` mentioned in the previous section
# This step allows the sonobuoy pod to read the test configurations.
$ kubectl create ns sonobuoy && \
  kubectl create cm l5d-config \
  -n sonobuoy \
  --from-file=config.yaml=/path/to/config.yaml

# Run the plugin
$ plugin=https://raw.githubusercontent.com/linkerd/linkerd2-conformance/master/sonobuoy/plugin.yaml

$ sonobuoy run \
  --plugin $plugin \
  --skip-preflight \
  --wait

# [Optional] Check the status of the pod
$ sonobuoy status

# Retrieve the test results
# This command downloads the tar ball containing the results
$ results=$(sonobuoy retrieve)

# Inspect the results
$ sonobuoy results $results --mode dump

# Clean up your cluster
$ sonobuoy delete --wait
```

### Running the tests locally

These commands assume a working [Go 1.14](https://golang.org/doc/go1.14) environment along with the [Linkerd2 CLI](https://linkerd.io/2/getting-started/#step-1-install-the-cli) and [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) (with a correctly configured `kubeconfig`) installed.

```bash
# clone this repository
$ git clone https://github.com/linkerd/linkerd2-conformance

# Navigate into project directory
$ cd linkerd2-conformance

# Use the convenient script to run `go run`
# $ bin/go-test [TEST NAME] [Test ID]

$ bin/go-test inject 01

# To run all the tests in the recomended order
$ bin/run-all
```

> Note: The `Test ID` argument only serves the purpose of ordering the JUnit reports for Sonobuoy to aggregate the results. While running a single test locally, any arbitrary number can be provided

On completion of the test(s), a `reports/` folder will be created at the root directory of the project, from where the JUnit reports of each individual test can be accessed.

## Adding new tests

This project makes use of [Ginkgo](https://github.com/onsi/ginkgo) paired with [Gomega](https://github.com/onsi/gomega) matcher library to describe tests and write assertions. Each of the tests can be found under the `tests/` folder, in its respective packages.

To add a new test, follow the steps below.

#### 1. Initial bootstrapping

Copy the `test_template/` folder into the `tests/` directory, and rename it to the name of the test that is being added. 

For example, let's assume a sample Linkerd2 feature called `l5dFeature`.

```bash
$ cp -r test_template/ tests/
$ mv tests/test_template tests/l5dFeature
$ mv tests/l5dFeature/template_test.go tests/l5dFeature/l5dFeature_test.go
```

#### 2. Adding tests

Under the `tests/l5dFeature` directory, you will find 3 files

- `l5dFeature_test.go` - the entry point for our `l5dFeature` test suite
- `specs.go` - the structure of our [Ginkgo](https://github.com/onsi/ginkgo) specs are defined here
- `tests.go` - [Gomega](https://github.com/onsi/gomega) assertions and test logic organized into separate functions, each corresponding to a spec defined in `specs.go`

Replace every instance of `<testname>` with the name of the test / package (in this case - `l5dFeature`). An example of what the file contents must look like is shown below.

```go
// l5dFeature_test.go

package l5dFeature

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(utils.BeforeSuiteCallback)

func TestL5dFeature(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "L5d feature")
}

var _ = AfterSuite(utils.AfterSuiteCallBack)

```

```go
// specs.go
package l5dFeature

import (
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("L5d Featuure", func() {
	h, c := utils.GetHelperAndConfig()

	ginkgo.It("can do this", testCanDoThis)
	ginkgo.It("can do that", testCanDoThat)
})

```

```go
// tests.go

package l5dFeature

func testCanDoThis() {
	// add assertions and test logic here
}

func testCanDoThat() {
	// add assertions and test logic here
}
```

#### 3. Running the test locally

Follow the instructions under [this section](https://github.com/mayankshah1607/linkerd2-conformance#running-the-tests-locally) to run the newly added test locally.

#### 4. Wiring up newly added test

Once the test has been written correctly and is working as expected, include it under `bin/run-all`.

For example

```bash
#!/bin/bash
set -eu

# Test IDs are important because thats the order in which sonobuoy will aggregate results

bin/go-test lifecycle 01

# Add tests here
bin/go-test inject 02
bin/go-test l5dFeature 03 # Newly added tests

# Should always run at the end
bin/go-test uninstall 03
```

<!-- refs -->
[logo]: https://user-images.githubusercontent.com/9226/33582867-3e646e02-d90c-11e7-85a2-2e238737e859.png
