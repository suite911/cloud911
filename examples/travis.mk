# Travis CI runs on Ubuntu, so it is safe to assume that GNU make is installed.
# This is a build script for GNU make

all:
	gmake -C 0-hello -f ../travis.mk example

example:
	go build

.PHONY: all example
