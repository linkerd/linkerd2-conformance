#!/bin/bash

# This script is to be run from inside the Sonobuoy pod

# curl -sL https://raw.githubusercontent.com/linkerd/website/master/run.linkerd.io/public/install | sh

/bin/bash -c "./conformance --ginkgo.reportFile=/tmp/results/report.xml"


tar czvf /tmp/results/results.tar.gz -C /tmp/results .

echo -n /tmp/results/results.tar.gz > /tmp/results/done
