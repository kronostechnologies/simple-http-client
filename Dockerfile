FROM --platform=$BUILDPLATFORM golang:1.17 AS builder
WORKDIR /go/src/github.com/kronostechnologies/simple-http-client/
COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -ldflags="-w -s" -o simple-http-client .
RUN echo "nobody:x:65534:65534:nobody:/:" > /tmp/passwd

FROM scratch
COPY --from=builder /go/src/github.com/kronostechnologies/simple-http-client/simple-http-client /bin/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /tmp/passwd /etc/passwd

USER 65534:65534
ENTRYPOINT ["/bin/simple-http-client"]
