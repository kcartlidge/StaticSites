echo Building Windows edition
env GOOS=windows GOARCH=amd64 go build -o builds/windows/StaticSites.exe

echo Building MacOS edition
env GOOS=darwin GOARCH=amd64 go build -o builds/macos/StaticSites

echo Building Linux edition
env GOOS=linux GOARCH=amd64 go build -o builds/linux/StaticSites

echo Copying into the sample static side builds folder
cp builds/windows/StaticSites.exe ../sites/staticsites.io/builds/windows/StaticSites.exe
cp builds/macos/StaticSites ../sites/staticsites.io/builds/macos/StaticSites
cp builds/linux/StaticSites ../sites/staticsites.io/builds/linux/StaticSites
