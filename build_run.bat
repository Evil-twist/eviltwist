@echo off
set GOARCH=amd64
echo Building...
go build -o .\build\exodus.exe -mod=vendor && cls && .\build\exodus.exe -p ./phishlets -t ./redirectors -developer -debug
