#!/usr/bin/env sh
# Simple DNS challenge exec solver.
# Use challtestsrv https://github.com/letsencrypt/pebble/tree/master/cmd/pebble-challtestsrv

set -e

case "$1" in
  "present")
    echo  "Present"
    payload="{\"host\":\"$2\", \"value\":\"$3\"}"
    # payload="{\"host\":\"test\", \"value\":\"test\"}"
    echo "payload=${payload}"
    # curl -s -X POST -d "${payload}" challtestsrv:8055/set-txt
    wget --quiet -O /dev/null --post-data "${payload}" http://challtestsrv:8055/set-txt
    ;;
  "cleanup")
    echo  "cleanup"
    payload="{\"host\":\"$2\"}"
    echo "payload=${payload}"
    # curl -s -X POST -d "${payload}" challtestsrv:8055/clear-txt
    wget --quiet -O /dev/null --post-data "${payload}" http://challtestsrv:8055/clear-txt
    ;;
  *)
    echo "OOPS"
    ;;
esac
