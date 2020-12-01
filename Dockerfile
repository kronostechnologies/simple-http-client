FROM golang:1.15 AS builder
RUN apt update ; apt install upx-ucl -y ; apt clean
WORKDIR /go/src/github.com/kronostechnologies/simple-http-client/
COPY * ./
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o simple-http-client . && upx --best simple-http-client

FROM scratch
COPY --from=builder /go/src/github.com/kronostechnologies/simple-http-client/simple-http-client /bin/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/bin/simple-http-client"]
