.POSIX:

gopath="$${GOPATH:-$${HOME}/go}"
xdg_config_home="$${XDG_CONFIG_HOME:-$${HOME}/.config}"
srv911=$(gopath)/src/github.com/amy911/srv911

GO=go
USER=$(xdg_config_home)/amy911/srv911/user

all: update build

build: .phony
	@if test -d $(USER); then cp -frv $(USER) $(srv911)/user ; fi
	$(GO) build

update: .phony
	$(GO) get -u github.com/amy911/srv911

.phony:
