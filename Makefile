VERSION = 0.1.0

RELEASE_PKG = ./cmd/goggles
INSTALL_PKG = $(RELEASE_PKG)

# Runs goggles and opens the logs.
#
# This is the default command.
run: | run.goggles logs
	@# Do Nothing
.PHONY: run.logs

# Runs gulp on the static assets.
gulp:
	@gulp
.PHONY: gulp

# Cleans any built artifacts.
clean:
	@rm -rf goggles.app
	@rm -f ~/Library/Logs/goggles.log
.PHONY: clean

# Builds goggles to the ./bin directory.
build: | clean gulp
	@mkdir -p bin/
	@go build -v -o bin/goggles $(INSTALL_PKG)
	@gallium-bundle bin/goggles
	@mkdir -p goggles.app/Contents/MacOS/static
	@cp -r ./_static/ goggles.app/Contents/MacOS/static
.PHONY: build

# Runs goggles.
run.goggles: | build
	@pkill goggles || true
	@open goggles.app
.PHONY: run

# Opens the logs.
logs: 
	@tail -100f ~/Library/Logs/goggles.log
.PHONY: logs

# Remote includes require 'mmake' 
# github.com/tj/mmake
include github.com/KyleBanks/make/go/sanity
include github.com/KyleBanks/make/go/release
include github.com/KyleBanks/make/git/precommit