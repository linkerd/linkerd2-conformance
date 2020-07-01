FROM golang:1.14


WORKDIR /linkerd2-conformance

COPY bin/ bin/
COPY sonobuoy/ sonobuoy/
COPY tests/ tests/
COPY utils/ utils/
COPY go.mod .
COPY go.sum .
# Build the test binary
RUN go test -i ./tests/...

RUN apt update && \ 
  apt upgrade -y && \
  apt install curl -y

# install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin

