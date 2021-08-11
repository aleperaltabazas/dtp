PWD=$(shell pwd)
EXE_NAME?=dtp

build:
	go build -o $(PWD)/$(EXE_NAME) github.com/aleperaltabazas/dtp/main

clean:
	rm $(EXE_NAME)

install: clean build
	mkdir -p ~/.local/bin
	cp $(EXE_NAME) ~/.local/bin
