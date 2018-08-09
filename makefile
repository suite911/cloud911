.POSIX:

gopath="$${GOPATH:-$${HOME}/go}"
xdg_config_home="$${XDG_CONFIG_HOME:-$${HOME}/.config}"
srv911=$(gopath)/src/github.com/amy911/srv911

GO=go
SECRET=$(xdg_config_home)/amy911/srv911/secret

all: update build

build: .phony
	@if test -e $(srv911)/secret; then rm -frv $(srv911)/secret ; fi
	@if test -d $(SECRET); then cp -rv $(SECRET) $(srv911)/secret ; fi
	@if test -n -e $(srv911)/secret; then cp -rv $(srv911)/dfl/secret $(srv911)/secret ; fi
	$(GO) build

update: .phony
	$(GO) get -u github.com/amy911/srv911

.phony:
