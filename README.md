# goggles

[![GoDoc](https://godoc.org/github.com/KyleBanks/goggles?status.svg)](https://godoc.org/github.com/KyleBanks/goggles)&nbsp; 
[![Build Status](https://travis-ci.org/KyleBanks/goggles.svg?branch=master)](https://travis-ci.org/KyleBanks/goggles)&nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/goggles)](https://goreportcard.com/report/github.com/KyleBanks/goggles)&nbsp;
[![Coverage Status](https://coveralls.io/repos/github/KyleBanks/goggles/badge.svg?branch=master)](https://coveralls.io/github/KyleBanks/goggles?branch=master)

🔭  Goggles is a GUI for your $GOPATH.

![Goggles Demo](./demo.gif)

## Install

### Stable

Grab the latest release from the [Releases](https://github.com/KyleBanks/goggles/releases) page. 

**Note:** It is currently a known issue that Goggles can only be run via the command-line with `open Goggles.app`, simply double-clicking `Goggles.app` has issues finding your $GOPATH.

### From Source

In order to build and run Goggles, there are a few steps you'll need to take:

1. `go get github.com/KyleBanks/goggles/...`
2. Install [Gallium](https://github.com/alexflint/gallium), in order to bundle the `.app`.
3. Install [npm](https://www.npmjs.com/) and [Gulp](http://gulpjs.com/), in order to build the front-end assets.
4. Run `make` to build and launch the application.

**Note:** Goggles is currently only available for Mac OS. If you'd like to see Goggles available on additional platforms, I encourage you to help contribute to the [Gallium](https://github.com/alexflint/gallium) project.

## Contributing

Contributions to Goggles are very welcome! In order to contribute, either open a new issue for discussion prior to making changes, or comment on an existing ticket indicating that you'd like to take it.

## Author

Goggles was developed by [Kyle Banks](https://twitter.com/kylewbanks).

## Credits

The [Gopher](./_static/img) loading images were created by [Ashley McNamara](https://twitter.com/ashleymcnamara) and inspired by [Renee French](http://reneefrench.blogspot.co.uk/).

![Gopher](./_static/img/loader-1.png)

## License

Goggles is available under the [Apache 2.0](./LICENSE) license.
