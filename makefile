.POSIX:

gopath="$${GOPATH:-$${HOME}/go}"
xdg_config_home="$${XDG_CONFIG_HOME:-$${HOME}/.config}"

GO=go
SECRET=$(xdg_config_home)/amy911/srv911/secret

all: update build

build: .phony
	cd $(gopath)/src/github.com/amy911/srv911
	@cp -v $(SECRET) ./
	$(GO) build
	@cd -

update: .phony
	$(GO) get -u github.com/amy911/srv911

.phony:
