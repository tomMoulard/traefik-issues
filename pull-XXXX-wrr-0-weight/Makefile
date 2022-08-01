LEN := 10000
NUMBERS := $(shell seq 1 ${LEN})
JOBS := $(addprefix job,${NUMBERS})

all: clean dos report

dos: ${JOBS}
${JOBS}: job%: ; @curl http://localhost | grep Name >> file.txt

report:
	cat file.txt | sort | uniq -c

clean:
	$(RM) file.txt
