########################################################################################

# This Makefile generated by GoMakeGen 0.5.0 using next command:
# gomakegen --metalinter .

########################################################################################

.PHONY = fmt all clean deps metalinter

########################################################################################

all: knfgen

knfgen:
	go build knfgen.go

deps:
	git config --global http.https://pkg.re.followRedirects true
	go get -d -v pkg.re/essentialkaos/ek.v8

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;

metalinter:
	test -s $(GOPATH)/bin/gometalinter || (go get -u github.com/alecthomas/gometalinter ; $(GOPATH)/bin/gometalinter --install)
	$(GOPATH)/bin/gometalinter --deadline 30s

clean:
	rm -f knfgen

########################################################################################
