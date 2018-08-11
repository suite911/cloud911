.POSIX:

gopath="$${GOPATH:-$${HOME}/go}"
xdg_config_home="$${XDG_CONFIG_HOME:-$${HOME}/.config}"
cloud911=$(gopath)/src/github.com/amy911/cloud911

GO=go
USER=$(xdg_config_home)/amy911/cloud911/user

all: update build

build: .phony
	@if test -d $(USER); then cp -frv $(USER) $(cloud911)/user ; fi
	$(GO) build

travis: .phony
	gmake -C examples -f travis.mk

update: .phony
	$(GO) get -u github.com/amy911/cloud911

.phony:
