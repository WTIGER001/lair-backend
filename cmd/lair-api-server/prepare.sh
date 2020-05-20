#!/bin/bash
 
# CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .
CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' .
# cp $GOPATH/bin/lair-api-server lair-api-server