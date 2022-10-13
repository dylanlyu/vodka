SCRIPTS = $(shell pwd)/scripts
PROJECT = $(shell pwd)
GO = $(shell which go)
OUTPUTS = $(shell pwd)/deploy
TAG ?= debug

authority:
	GOOS=linux GOARCH=amd64 $(GO) build -tags $(TAG) -o $(OUTPUTS)/authority cmd/lambda/authority.go
	zip -D -j -r $(OUTPUTS)/authority.zip $(OUTPUTS)/authority

clean:
	rm -rf $(OUTPUTS)

task:
	make clean
	make authority

setup:
	cp $(PROJECT)/config/config.go.example $(PROJECT)/config/debug_config.go
	cp $(PROJECT)/config/config.go.example $(PROJECT)/config/production_config.go
	cp $(PROJECT)/air.example $(PROJECT)/.air.toml

air:
	air

migration:
	go run -tags $(TAG) $(PROJECT)/tools/migration/migration.go

ssh:
	go run -tags $(TAG) $(PROJECT)/tools/ssh/ssh.go

changeLog:
	git-chglog > ./changeLog.md