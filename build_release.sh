#!/usr/bin/env bash

PROJECT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}"  )" && pwd  )"
PROJECT_NAME="$(basename "${PROJECT_PATH}" )"

RELEASE_VERSION="$(go run ${PROJECT_NAME}/cli.go -v)"
RELEASE_FILE=${PROJECT_PATH}/${PROJECT_NAME}_${RELEASE_VERSION}.zip
BUILD_DIR=${PROJECT_PATH}/build

mkdir -p ${BUILD_DIR}

GOOS=darwin GOARCH=386 go build -o ${BUILD_DIR}/${PROJECT_NAME}_darwin_386 ${PROJECT_NAME}/cli.go
GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_darwin_amd64 ${PROJECT_NAME}/cli.go
GOOS=dragonfly GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_dragonfly_amd64 ${PROJECT_NAME}/cli.go
GOOS=freebsd GOARCH=386 go build -o ${BUILD_DIR}/${PROJECT_NAME}_freebsd_386 ${PROJECT_NAME}/cli.go
GOOS=freebsd GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_freebsd_amd64 ${PROJECT_NAME}/cli.go
GOOS=freebsd GOARCH=arm go build -o ${BUILD_DIR}/${PROJECT_NAME}_freebsd_arm ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=386 go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_386 ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_amd64 ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=arm go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_arm ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=arm64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_arm64 ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=ppc64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_ppc64 ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=ppc64le go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_ppc64le ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=mips64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_mips64 ${PROJECT_NAME}/cli.go
GOOS=linux GOARCH=mips64le go build -o ${BUILD_DIR}/${PROJECT_NAME}_linux_mips64le ${PROJECT_NAME}/cli.go
GOOS=netbsd GOARCH=386 go build -o ${BUILD_DIR}/${PROJECT_NAME}_netbsd_386 ${PROJECT_NAME}/cli.go
GOOS=netbsd GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_netbsd_amd64 ${PROJECT_NAME}/cli.go
GOOS=netbsd GOARCH=arm go build -o ${BUILD_DIR}/${PROJECT_NAME}_netbsd_arm ${PROJECT_NAME}/cli.go
GOOS=openbsd GOARCH=386 go build -o ${BUILD_DIR}/${PROJECT_NAME}_openbsd_386 ${PROJECT_NAME}/cli.go
GOOS=openbsd GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_openbsd_amd64 ${PROJECT_NAME}/cli.go
GOOS=openbsd GOARCH=arm go build -o ${BUILD_DIR}/${PROJECT_NAME}_openbsd_arm ${PROJECT_NAME}/cli.go
GOOS=plan9 GOARCH=386 go build -o ${BUILD_DIR}/${PROJECT_NAME}_plan9_386 ${PROJECT_NAME}/cli.go
GOOS=plan9 GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_plan9_amd64 ${PROJECT_NAME}/cli.go
GOOS=solaris GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_solaris_amd64 ${PROJECT_NAME}/cli.go
GOOS=windows GOARCH=386 go build -o ${BUILD_DIR}/${PROJECT_NAME}_windows_386.exe ${PROJECT_NAME}/cli.go
GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${PROJECT_NAME}_windows_amd64.exe ${PROJECT_NAME}/cli.go

cd ${BUILD_DIR} && zip -r ${RELEASE_FILE} *
cd ${PROJECT_PATH}
rm -rf ${BUILD_DIR}
echo ${RELEASE_FILE}
