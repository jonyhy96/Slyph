#! /usr/bin/make -f
targetImagename = slyph

.PHONY: all
all: package

.PHONY: test
test: build

.PHONY: package
package: build createImage

.PHONY: build
build: 
	go build -o Slyph main.go

.PHONY: createImage	
createImage:
	docker build -t $(targetImagename) .

.PHONY: clean	
clean:
    rm -f Slyph 
	docker rmi $(docker images|grep slyph|awk '{print $1}') 