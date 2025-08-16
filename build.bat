@echo off
go build -ldflags "-H windowsgui -s -w" -o termnia.exe ./cmd/termnia