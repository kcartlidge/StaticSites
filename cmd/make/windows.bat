@echo off
@echo Building Linux edition
@set GOOS=linux
@set GOARCH=amd64
@go build -o builds/linux/StaticSites

@echo Building MacOS edition
@set GOOS=darwin
@set GOARCH=amd64
@go build -o builds/macos/StaticSites

@echo Building Windows edition
@set GOOS=windows
@set GOARCH=amd64
@go build -o builds/windows/StaticSites.exe
