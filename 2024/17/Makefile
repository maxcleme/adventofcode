ANTLR_VERSION=4.13.2
UID=$(shell id -u)
GID=$(shell id -g)
PWD=$(shell pwd)

antlr-image:
	docker build --build-arg uid=${UID} --build-arg gid=${GID} --tag antlr:${ANTLR_VERSION} https://github.com/antlr/antlr4.git#${ANTLR_VERSION}:docker/

generate-go-antlr: antlr-image
	docker run -v ${PWD}:/work antlr:4.13.2 -Dlanguage=Go -o parser AOC.g4