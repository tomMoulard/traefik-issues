#!/bin/bash

FILE=$1

upload() {
	MC_ENV=MC_HOST_MINIO=https://minioadmin:minioadmin@minio.localhost \
	mc cp ./${FILE} test-bucket/$1
}

for i in {1..200}; do
	upload $i &
done
