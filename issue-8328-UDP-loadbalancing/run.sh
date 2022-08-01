#!/bin/bash

func () {
    echo WHO | nc -u localhost 8000
}

export -f func

parallel --timeout 1s func ::: {00..100} 2>/dev/null \
    | grep Name \
    | sort \
    | uniq -c

