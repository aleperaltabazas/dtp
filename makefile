PWD=$(shell pwd)
EXE_NAME?=dtp

build:
	go build -o $(PWD)/$(EXE_NAME) github.com/aleperaltabazas/dtp/main
