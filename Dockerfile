FROM golang:1-buster AS builder
ENV GO111MODULE on
RUN mkdir /src
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /src
WORKDIR /src
RUN make build

FROM debian:buster
COPY --from=builder /src/cli-template /usr/local/bin
ENTRYPOINT ["/usr/local/bin/cli-template"]
