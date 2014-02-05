VERSION = `awk '/VERSION/ { gsub("\"", ""); print $$NF }' $(CURDIR)/lib/version.go`
DIST_DIR = $(CURDIR)/dist

#
# Packaging options
#
XC_ARCH ?= "386 amd64 arm"
XC_OS   ?= "darwin linux freebsd openbsd"

build: clean
	go build

clean: 
	rm -rf gill

package:
	@(echo $(VERSION))
	@(go get github.com/laher/goxc)
	@($(GOPATH)/bin/goxc -arch=$(XC_ARCH) -os=$(XC_OS) -pv=$(VERSION) -d=$(DIST_DIR))

.PHONY: install
