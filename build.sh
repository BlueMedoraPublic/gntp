#!/bin/sh

VERSION=`grep 'const version' gntp.go | awk '{print $5}' | tr -d '"'`
if [ -z "$VERSION" ]
then
      echo "Failed to get version from gntp.go"
      exit 1
else
      echo "Building gntp ${VERSION}"
fi

mkdir -p artifacts

docker build \
    --no-cache \
    --build-arg version=${VERSION} \
    -t gntp:${VERSION} .

docker create -ti --name artifacts gntp:${VERSION} bash && \
    docker cp artifacts:/gntp_linux_amd64.zip artifacts/gntp_linux_amd64.zip && \
    docker cp artifacts:/gntp_darwin_amd64.zip artifacts/gntp_darwin_amd64.zip && \
    docker cp artifacts:/gntp_windows_amd64.zip artifacts/gntp_windows_amd64.zip && \
    docker cp artifacts:/SHA256SUMS artifacts/SHA256SUMS

# cleanup
docker rm -fv artifacts &> /dev/null
