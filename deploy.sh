#!/bin/sh
git pull
export PATH=$PATH:/usr/local/go/bin
go build -v
kill -9 $(ps aux | grep '[a]pp-paykeen' | awk '{print $2}')
nohup ./app-paykeen&