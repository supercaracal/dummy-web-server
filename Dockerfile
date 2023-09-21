FROM golang:1.20 as builder
WORKDIR /go/src/app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -tags timetzdata -o server

# https://github.com/GoogleContainerTools/distroless
# https://console.cloud.google.com/gcr/images/distroless/GLOBAL
FROM gcr.io/distroless/static-debian12:nonroot-amd64
COPY --from=builder /go/src/app/server /usr/local/bin/dummy-web-server
ENTRYPOINT ["/usr/local/bin/dummy-web-server"]
