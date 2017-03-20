VERSION = 0.1.0

RELEASE_PKG = ./cmd/goggles
INSTALL_PKG = $(RELEASE_PKG)

APP_FOLDER = bin/goggles.app

# Runs goggles and opens the logs.
#
# This is the default command.
run: | run.goggles logs
	@# Do Nothing
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
	@rm -f ~/Library/Logs/goggles.log
.PHONY: clean

# Builds goggles to the ./bin directory.
build: | clean gulp
	@mkdir -p bin/
	@go build -v -o bin/goggles $(INSTALL_PKG)
	@gallium-bundle bin/goggles --output $(APP_FOLDER)
	@mkdir -p $(APP_FOLDER)/Contents/MacOS/static
	@cp -r ./_static/ $(APP_FOLDER)/Contents/MacOS/static
.PHONY: build

# Runs goggles.
run.goggles: | build
	@pkill goggles || true
	@open $(APP_FOLDER)
.PHONY: run

# Opens the logs.
logs: 
	@tail -100f ~/Library/Logs/goggles.log
.PHONY: logs

# Remote includes require 'mmake' 
# github.com/tj/mmake
include github.com/KyleBanks/make/go/sanity
include github.com/KyleBanks/make/git/precommit