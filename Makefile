VERSION = 0.1.1

INSTALL_PKG = ./cmd/goggles

BIN = ./bin
APP_NAME = Goggles.app
APP_FOLDER = $(BIN)/$(APP_NAME)
APP_STATIC_FOLDER = $(APP_FOLDER)/Contents/MacOS/static
LOG_FILE = ~/Library/Logs/goggles.log

BUNDLE_ID = com.kylewbanks.goggles
BUNDLE_NAME = Goggles

# Runs goggles and opens the logs.
#
# This is the default command.
run: | run.goggles
	@tail -100f $(LOG_FILE)
.PHONY: run.logs

# Runs gulp on the static assets.
gulp:
	@cd _static ; \
	npm install ; \
	gulp 
.PHONY: gulp

# Cleans any built artifacts.
clean:
	@rm -rf $(BIN)
	@rm -f $(LOG_FILE)
.PHONY: clean

# Builds goggles to the ./bin directory.
build: | clean gulp
	@mkdir -p bin/
	@go build -v -o bin/goggles $(INSTALL_PKG)
	@gallium-bundle bin/goggles \
		--output $(APP_FOLDER) \
		--identifier $(BUNDLE_ID) \
		--name $(BUNDLE_NAME)
	@mkdir -p $(APP_STATIC_FOLDER)
	@cp -a ./_static/. $(APP_STATIC_FOLDER)
	@rm -rf $(APP_STATIC_FOLDER)/node_modules
.PHONY: build

# Builds a release bundle.
release: | build 
	@cd $(BIN) ; \
	zip -r -y goggles.$(VERSION).zip $(APP_NAME)
.PHONY: release

# Runs the goggles application.
run.goggles: | build
	@pkill Goggles || true
	@open $(APP_FOLDER)
.PHONY: run

# Runs test cases in Docker.
test.docker:
	@docker build -t goggles-test .
	@docker run -it goggles-test
.PHONY: test.docker

# Runs test suit, vet, golint, and fmt.
sanity:
	@echo "---------------- TEST ----------------"
	@go list ./... | grep -v vendor/ | xargs go test --cover 

	@echo "---------------- VET ----------------"
	@go list ./... | grep -v vendor/ | xargs go vet 

	@echo "---------------- LINT ----------------"
	@go list ./... | grep -v vendor/ | xargs golint

	@echo "---------------- FMT ----------------"
	@go list ./... | grep -v vendor/ | xargs go fmt
.PHONY: sanity

# Installs a pre-commit Git hook that executes "make sanity" prior to commiting.
precommit:
	@echo "make sanity" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
.PHONY: precommit
