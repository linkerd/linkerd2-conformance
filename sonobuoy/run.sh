#!/bin/bash

# This script is to be run from inside the Sonobuoy pod

/bin/bash -c "./conformance --ginkgo.reportFile=/tmp/results/report.xml"


tar czvf /tmp/results/results.tar.gz -C /tmp/results .

echo -n /tmp/results/results.tar.gz > /tmp/results/done
