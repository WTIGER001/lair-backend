FROM scratch

COPY lair-api-server /lair-api-server
COPY cert.pem /cert.pem
COPY key.pem /key.pem

CMD ["/lair-api-server", "--tls-certificate", "cert.pem", "--tls-key", "key.pem", "--port", "80", "--tls-port", "443", "--host","0.0.0.0"]
# CMD ["/bin/ls"]