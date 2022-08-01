FILE=trash_file.bin

setup:
	${MC_ENV} mc mb test-bucket --ignore-existing
	head -c 1G </dev/urandom > ${FILE}

test:
	./test.sh ${FILE}

clean:
	$(RM) -r ./volumes/minio/*
	$(RM) -r ./volumes/minio/.*
	$(RM) ${FILE}
