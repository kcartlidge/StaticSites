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

@echo Copying into the sample static side builds folder
@copy builds\windows\StaticSites.exe ..\sites\staticsites.io\builds\windows\StaticSites.exe
@copy builds\macos\StaticSites ..\sites\staticsites.io\builds\macos\StaticSites
@copy builds\linux\StaticSites ..\sites\staticsites.io\builds\linux\StaticSites
