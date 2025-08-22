@echo off
setlocal

set APP=termnia
set SRC=./cmd/termnia
set BIN=bin

if not exist %BIN% mkdir %BIN%

echo ==== Building Windows (amd64) ====
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-H windowsgui -s -w" -o %BIN%\%APP%-windows-amd64.exe %SRC%

echo ==== Building Linux (amd64) ====
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -o %BIN%\%APP%-linux-amd64 %SRC%

echo ==== Building macOS (darwin-arm64) ====
set GOOS=darwin
set GOARCH=arm64
go build -ldflags "-s -w" -o %BIN%\%APP%-darwin-arm64 %SRC%

echo ==== Build completed! Binaries are in %BIN% ====
endlocal
pause
