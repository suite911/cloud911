.POSIX:

gopath="$${GOPATH:-$${HOME}/go}"
xdg_config_home="$${XDG_CONFIG_HOME:-$${HOME}/.config}"
cloud911=$(gopath)/src/github.com/suite911/cloud911

GO=go
USER=$(xdg_config_home)/suite911/cloud911/user

all: update build

build: .phony
	@if test -d $(USER); then cp -frv $(USER) $(cloud911)/user ; fi
	$(GO) build

update: .phony
	$(GO) get -u github.com/suite911/cloud911

.phony:
