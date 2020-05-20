#!/bin/bash

$GOPATH/bin/lair-api-server --tls-certificate=cert.pem --tls-key=key.pem --host=0.0.0.0 --port=4201 --tls-port=4401