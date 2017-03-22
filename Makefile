VERSION = 0.1.0

INSTALL_PKG = ./cmd/goggles

APP_FOLDER = ./bin/goggles.app
APP_STATIC_FOLDER = $(APP_FOLDER)/Contents/MacOS/static
LOG_FILE = ~/Library/Logs/goggles.log

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
	@rm -rf $(APP_FOLDER)
	@rm -f $(LOG_FILE)
.PHONY: clean

# Builds goggles to the ./bin directory.
build: | clean gulp
	@mkdir -p bin/
	@go build -v -o bin/goggles $(INSTALL_PKG)
	@gallium-bundle bin/goggles --output $(APP_FOLDER)
	@mkdir -p $(APP_STATIC_FOLDER)
	@cp -r ./_static/ $(APP_STATIC_FOLDER)
	@rm -rf $(APP_STATIC_FOLDER)/node_modules
.PHONY: build

# Runs the goggles application.
run.goggles: | build
	@pkill goggles || true
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
