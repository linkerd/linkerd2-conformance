FROM golang:1.14 as build

COPY . /conformance
WORKDIR /conformance

# Build the test binary
RUN go test -c -o conformance

FROM debian:bullseye

RUN apt update && \ 
    apt upgrade -y && \
    apt install curl -y

# install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \
    ./kubectl /usr/local/bin

# Copy test binary
COPY --from=build /conformance/conformance /conformance

COPY ./sonobuoy/run.sh .
COPY ./testdata /testdata

