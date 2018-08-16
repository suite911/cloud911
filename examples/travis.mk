# Travis CI runs on Ubuntu, so it is safe to assume that GNU make is installed.
# This is a build script for GNU make

all:
	make -C 0-hello -f ../travis.mk example
	make -C 1-basic -f ../travis.mk example

example:
	go generate
	go build
	go test

.PHONY: all example
