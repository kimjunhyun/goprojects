#!/bin/sh
# Shows the missing certiicate portion of the rosettacode.org example
exec 3>&1 4>&2
exec >/dev/null 2>&1
PASS=$(date +%A%w%Y%m%d%H%M%S%B)$$${RANDOM}${RANDOM}
echo ${PASS}
openssl genrsa -des3 -passout pass:${PASS} -out server.key 1024
openssl req -new -key server.key -passin pass:${PASS} -out server.csr <<CERT
EA


n/a

127.0.0.1
${USER}@127.0.0.1


CERT
openssl rsa -in server.key -passin pass:${PASS} -out server.key
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

go build server.go && {
  ./server & server=$!
  trap "kill ${server}" 0
  go run client.go >&3
}
rm server.{key,crt,csr}
exit 0
