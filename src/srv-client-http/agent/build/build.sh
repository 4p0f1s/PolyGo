#!/bin/bash

# ANSI color codes
GREEN='\033[0;32m'
YELLOW='\033[0;33m'

echo -e "${YELLOW}[*] ${GREEN}Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o agentlin ../agent.go 

echo -e "${YELLOW}[*] ${GREEN}Building for Linux 32bits..."
GOOS=linux GOARCH=amd64 go build -o agentlin32 ../agent.go 

echo -e "${YELLOW}[*] ${GREEN}Building for Linux (ARM)..."
GOOS=linux GOARCH=arm64 go build -o agentlin_arm ../agent.go 

echo -e "${YELLOW}[*] ${GREEN}Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o agentwin.exe ../agent.go 

echo -e "${YELLOW}[*] ${GREEN}Building for Windows 32bits..."
GOOS=windows GOARCH=386 go build -o agentwin32.exe ../agent.go 

echo -e "${YELLOW}[*] ${GREEN}Building for Windows (ARM)..."
GOOS=windows GOARCH=arm64 go build -o agentwin_arm.exe ../agent.go 

echo -e "${YELLOW}[*] ${GREEN}Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o agentmacos ../agent.go 

echo -e "${YELLOW}[*] ${GREEN}Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o agentmacos_arm ../agent.go 