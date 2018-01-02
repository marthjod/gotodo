PACKAGES ?= $(shell go list ./... | grep -v /vendor/ | grep -v /tests)

.PHONY: all
#all: vet lint errcheck megacheck test
all: vet lint megacheck test

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:
	STATUS=0; for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || STATUS=1; done; exit $$STATUS

.PHONY: errcheck
errcheck:
	STATUS=0; for PKG in $(PACKAGES); do errcheck -ignoretests $$PKG || STATUS=1; done; exit $$STATUS

.PHONY: megacheck
megacheck:
	STATUS=0; for PKG in $(PACKAGES); do megacheck $$PKG || STATUS=1; done; exit $$STATUS

.PHONY: test
test:
	STATUS=0; for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || STATUS=1; done; exit $$STATUS

