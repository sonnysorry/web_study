#!/bin/sh
go run main.go
go install -ldflags "-X main.SHA1VER=`git rev-parse HEAD` -X main.BUILDTIME=`date -u +%Y-%m-%dT%H:%M:%S`"
sudo apps -http :5000