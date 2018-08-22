.POSIX:

DAYS=30
OPENSSL=openssl

all: .phony
	$(OPENSSL) req -new -newkey rsa:4096 -x509 -sha256 -days $(DAYS) -nodes -out cert.pem -keyout key.pem
	cp -fv cert.pem key.pem 0-hello/
	cp -fv cert.pem key.pem 1-basic/
	cp -fv cert.pem key.pem 2-amy/

.phony:
