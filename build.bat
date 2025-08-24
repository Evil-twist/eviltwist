@echo off
set GOARCH=amd64
echo Building...
go build -o .\build\exodus.exe -mod=vendor
