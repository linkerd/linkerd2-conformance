FROM mayankshah1607/linkerd2-conformance-base

RUN curl -LJO https://github.com/vmware-tanzu/sonobuoy/releases/download/v0.18.2/sonobuoy_0.18.2_linux_amd64.tar.gz \
    && tar -xvf sonobuoy_0.18.2_linux_amd64.tar.gz

RUN mv ./sonobuoy /usr/local/bin

WORKDIR /plugin

CMD sonobuoy run --plugin linkerd2-conformance.yaml --wait; tail -f /dev/null
