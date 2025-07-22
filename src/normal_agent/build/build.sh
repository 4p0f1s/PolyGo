#!/bin/bash

# ANSI color codes
GREEN='\033[0;32m'
YELLOW='\033[0;33m'

echo -e "${YELLOW}[*] ${GREEN}Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o revlin ../main.go ../unix.go

echo -e "${YELLOW}[*] ${GREEN}Building for Linux 32bits..."
GOOS=linux GOARCH=amd64 go build -o revlin32 ../main.go ../unix.go

echo -e "${YELLOW}[*] ${GREEN}Building for Linux (ARM)..."
GOOS=linux GOARCH=arm64 go build -o revlin_arm ../main.go ../unix.go

echo -e "${YELLOW}[*] ${GREEN}Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o revwin.exe ../main.go ../windows.go

echo -e "${YELLOW}[*] ${GREEN}Building for Windows 32bits..."
GOOS=windows GOARCH=386 go build -o revwin32.exe ../main.go ../windows.go

echo -e "${YELLOW}[*] ${GREEN}Building for Windows (ARM)..."
GOOS=windows GOARCH=arm64 go build -o revwin_arm.exe ../main.go ../windows.go

echo -e "${YELLOW}[*] ${GREEN}Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o revmacos ../main.go ../unix.go

echo -e "${YELLOW}[*] ${GREEN}Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o revmacos_arm ../main.go ../unix.go