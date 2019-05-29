# staging environment retrieves dependencies and compiles
#
FROM golang:1.12 AS stage

WORKDIR /build/src/gntp

ARG version

ADD . /build/src/gntp

RUN \
    go get github.com/mitchellh/gox && \
    go get github.com/beevik/ntp

RUN \
    go test ./...

RUN env CGO_ENABLED=0 $GOPATH/bin/gox -arch=amd64 -os='!netbsd !openbsd !freebsd'


# build the release with an image that includes zip and sha256sum
#
FROM debian:stable

WORKDIR /

ARG version

RUN apt-get update >> /dev/null && apt-get install zip -y

COPY --from=stage /build/src/gntp/gntp_linux_amd64 .
COPY --from=stage /build/src/gntp/gntp_darwin_amd64 .
COPY --from=stage /build/src/gntp/gntp_windows_amd64.exe .


RUN mv gntp_linux_amd64 gntp && zip gntp_linux_amd64.zip gntp
RUN mv gntp_darwin_amd64 gntp && zip gntp_darwin_amd64.zip gntp
RUN mv gntp_windows_amd64.exe gntp.exe && zip gntp_windows_amd64.zip gntp.exe


RUN ls | grep -v '.go' | grep 'gntp_' | xargs -n1 sha256sum >> SHA256SUMS
