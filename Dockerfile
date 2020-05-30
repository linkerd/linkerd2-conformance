FROM golang:1.14 as build

COPY . /conformance
WORKDIR /conformance

RUN go test -c -o conformance

FROM mayankshah1607/linkerd2-conformance-base 

RUN curl -sL https://run.linkerd.io/install | sh
COPY --from=build /conformance/conformance /conformance

CMD ["bash", "-c","/conformance"]

