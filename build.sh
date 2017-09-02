#!/bin/bash -eux

function build() {
    local version=$1
    local platform=$2
    local arch=$3

    GOOS=$platform GOARCH=$arch go build \
        -ldflags "-X main.Version=${version}" \
        -o go-exp-${platform}-${arch}
}

function main() {
    if [[ $# -ne 1 ]]; then
        echo "usage: ${0} <version>"
        exit 1
    fi
    local version=$1
    local platforms="linux darwin windows"
    for platform in $platforms; do
        echo calling to build for $platform
        build $version $platform "amd64"
        #Make windows binary executable
        if [[ $platform == "windows" ]] ; then
          mv go-exp-${platform}-amd64 go-exp-${platform}-amd64.exe
        fi
    done
}

main "$@"
