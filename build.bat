@echo off
go build -ldflags "-w -s" -o release/frpc.exe frpc.go
go build -ldflags "-w -s" -o release/frps.exe frps.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/frpc32.exe frpc.go
go build -ldflags "-w -s" -o release/frps32.exe frps.go

@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/frpc frpc.go
go build -ldflags "-w -s" -o release/frps frps.go
