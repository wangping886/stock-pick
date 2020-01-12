#!/usr/bin/env bash

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o stock-pick_linux main.go;

ox:
	go build -o stock-pick_ox  main.go;
