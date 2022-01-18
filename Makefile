SCRIPTS = $(shell pwd)/scripts
PROJECT = $(shell pwd)
GO = $(shell which go)
OUTPUTS = $(shell pwd)/deploy

authority:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(OUTPUTS)/authority cmd/lambda/authority.go
	zip -D -j -r $(OUTPUTS)/authority.zip $(OUTPUTS)/authority

clean:
	rm -rf $(OUTPUTS)

task:
	make clean
	make accounts

setup:
	cp $(SCRIPTS)/environment.example $(SCRIPTS)/debug.sh
	cp $(SCRIPTS)/environment.example $(SCRIPTS)/testing.sh
	cp $(SCRIPTS)/environment.example $(SCRIPTS)/production.sh
	cp $(SCRIPTS)/migration.example $(SCRIPTS)/migrationDebug.sh
	cp $(SCRIPTS)/migration.example $(SCRIPTS)/migrationTesting.sh
	cp $(SCRIPTS)/migration.example $(SCRIPTS)/migrationProduction.sh
	cp $(PROJECT)/air.example $(PROJECT)/.air.toml

debug:
	chmod a+x $(SCRIPTS)/debug.sh
	$(SCRIPTS)/debug.sh

testing:
	chmod a+x $(SCRIPTS)/testing.sh
	$(SCRIPTS)/testing.sh

production:
	chmod a+x $(SCRIPTS)/production.sh
	$(SCRIPTS)/production.sh

migrationDebug:
	chmod a+x $(SCRIPTS)/migrationDebug.sh
	$(SCRIPTS)/migrationDebug.sh

migrationTesting:
	chmod a+x $(SCRIPTS)/migrationTesting.sh
	$(SCRIPTS)/migrationTesting.sh

migrationProduction:
	chmod a+x $(SCRIPTS)/migrationProduction.sh
	$(SCRIPTS)/migrationProduction.sh

changeLog:
	git-chglog > ./README.md