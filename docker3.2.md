# MakeFile

export tag=v1.0
root:
	export ROOT=github.com/cncamp/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t cncamp/httpserver:${tag} .

push: release
	echo "pushing cncamp/httpserver"
	docker push cncamp/httpserver:v1.0

# Dockerfile

FROM ubuntu
LABEL multi.label1="value1" multi.label2="value2" other="value3"
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver
