#!/bin/sh
for f in client.go server.go authenticated.sh; do
  [ -f ${f} ] || curl -ksO https://gist.github.com/raw/4375261/${f} &
done
wait
sh authenticated.sh
exit 0
