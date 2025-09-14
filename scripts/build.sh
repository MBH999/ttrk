#!/bin/bash

# Build script for ttrk

set -e

# Variables
BINARY_NAME="ttrk"
BUILD_DIR="bin"
MAIN_PATH="./cmd/ttrk"

# Get version information
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DATE=$(date +%Y-%m-%d:%H:%M:%S)
GIT_COMMIT=$(git rev-parse HEAD 2>/dev/null || echo "unknown")

# LDFLAGS for version information
LDFLAGS="-X github.com/MBH999/ttrk/pkg/version.GitCommit=${GIT_COMMIT} -X github.com/MBH999/ttrk/pkg/version.BuildDate=${BUILD_DATE}"

echo "Building ${BINARY_NAME}..."
echo "Version: ${VERSION}"
echo "Build Date: ${BUILD_DATE}"
echo "Git Commit: ${GIT_COMMIT}"

# Create build directory
mkdir -p ${BUILD_DIR}

# Build the binary
go build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${BINARY_NAME} ${MAIN_PATH}

echo "Build complete: ${BUILD_DIR}/${BINARY_NAME}"