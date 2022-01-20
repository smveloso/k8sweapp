FROM golang:1.17.6
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY go.mod .
#COPY go.sum .
RUN go mod download all
COPY Makefile .
COPY cmd cmd
#COPY pkg pkg
RUN make build
FROM debian:buster
WORKDIR /
COPY --from=0 /go/src/app/k8swebapp .
ENTRYPOINT ["./k8swebapp"]
