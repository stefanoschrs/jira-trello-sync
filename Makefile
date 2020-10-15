.PHONY: run build

entr := $(shell which entr)
upx := $(shell which upx)

run:
ifdef entr
	find . -name '*.go' 2>&1 | entr -r go run .
else
	go run .
endif

build:
	go build .
ifdef upx
	upx jira-trello-sync
endif
