VERSION = 0.3.0

BIN = ./bin

INSTALL_PKG = ./cmd/goggles
RELEASE_PLATFORMS = darwin/386 darwin/amd64 linux/386 linux/amd64 linux/arm windows/386 windows/amd64

APP_INSTALL_PKG = ./cmd/goggles-app
APP_NAME = Goggles.app
APP_FOLDER = $(BIN)/$(APP_NAME)
APP_LOG_FILE = ~/Library/Logs/goggles.log
APP_BUNDLE_ID = com.kylewbanks.goggles
APP_BUNDLE_NAME = Goggles

# Runs Goggles within a web browser.
#
# This is the default command.
browser: | install
	@goggles
.PHONY: browser

# Runs Goggles as a standalone app and opens the logs.
app: | build.app
	@pkill Goggles || true
	@open $(APP_FOLDER)
	@tail -100f $(APP_LOG_FILE)
.PHONY: app

# Cleans, lints, and generates static assets.
assets:
	@cd static ; \
	npm install ; \
	gulp 

	@go-bindata-assetfs -ignore=node_modules -pkg assets static/... ; \
	mv bindata_assetfs.go server/assets/
.PHONY: assets

# Cleans any built artifacts.
clean:
	@rm -rf $(BIN)
	@rm -f $(APP_LOG_FILE)
.PHONY: clean

# Builds and installs Goggles as a browser application.
install: | clean assets	
	@go install -v $(INSTALL_PKG)
.PHONY: install

# Builds Goggles as a standalone app to the ./bin directory.
build.app: | clean assets
	@mkdir -p $(BIN)
	@go build -v -o $(BIN)/goggles-app $(APP_INSTALL_PKG)
	@gallium-bundle $(BIN)/goggles-app \
		--output $(APP_FOLDER) \
		--identifier $(APP_BUNDLE_ID) \
		--name $(APP_BUNDLE_NAME)
.PHONY: build.app

# Builds a release bundle for the standalone application, and
# binaries for the browser application.
release: | sanity build.app
	@cd $(BIN) ; \
	zip -r -y goggles.osx.app.$(VERSION).zip $(APP_NAME)

	@gox -osarch="$(strip $(RELEASE_PLATFORMS))" \
         -output "bin/{{.Dir}}_$(VERSION)_{{.OS}}_{{.Arch}}" $(INSTALL_PKG)
.PHONY: release

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
	@go list ./... | grep -v -e vendor/ -e server/assets | xargs golint

	@echo "---------------- FMT ----------------"
	@go list ./... | grep -v vendor/ | xargs go fmt
.PHONY: sanity

# Installs a pre-commit Git hook that executes "make sanity" prior to commiting.
precommit:
	@echo "make sanity" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
.PHONY: precommit
