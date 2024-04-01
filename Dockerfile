FROM alpine:latest
ADD ./build/release /go/bin/release
ENTRYPOINT /go/bin/release
