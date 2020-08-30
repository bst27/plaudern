#!/usr/bin/env bash
#
# This script builds the application from source for multiple platforms.

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Delete the old dir
echo "==> Removing old directory..."
rm -rf bin/*

# Create output dir
mkdir -p bin/
mkdir -p "bin/darwin-amd64"
mkdir -p "bin/linux-amd64"
mkdir -p "bin/linux-arm"
mkdir -p "bin/linux-arm64"
mkdir -p "bin/windows-amd64"

# Build Go project
echo "==> Building..."
env GOOS=darwin GOARCH=amd64 go build -v -o "bin/darwin-amd64" "$DIR/cmd/plaudern/"
env GOOS=linux GOARCH=amd64 go build -v -o "bin/linux-amd64" "$DIR/cmd/plaudern/"
env GOOS=linux GOARCH=arm go build -v -o "bin/linux-arm" "$DIR/cmd/plaudern/"
env GOOS=linux GOARCH=arm64 go build -v -o "bin/linux-arm64" "$DIR/cmd/plaudern/"
env GOOS=windows GOARCH=amd64 go build -v -o "bin/windows-amd64" "$DIR/cmd/plaudern/"

# Build Angular project
cd "web"
npm install
ng build --prod --base-href "/manage/"

# Copy Angular dist files
cd "$DIR"
cp -r "web/dist/web" "bin/darwin-amd64/"
cp -r "web/dist/web" "bin/linux-amd64/"
cp -r "web/dist/web" "bin/linux-arm"
cp -r "web/dist/web" "bin/linux-arm64"
cp -r "web/dist/web" "bin/windows-amd64"