#!/usr/bin/bash
env GOOS=linux GOARCH=amd64; go build -o licenser_linux_amd64
env GOOS=linux GOARCH=arm64; go build -o licenser_linux_arm64
env GOOS=darwin GOARCH=amd64; go build -o licenser_macos_amd64
env GOOS=darwin GOARCH=arm64; go build -o licenser_macos_arm64
env GOOS=windows GOARCH=amd64; go build -o licenser_windows_amd64.exe