# Run tests in alpine linux to ensure that the MacOS-only portions of the application do 
# not break testing on Travis.CI.
FROM golang:1.8.0-alpine

MAINTAINER Kyle Banks 

RUN mkdir -p $GOPATH/src/github.com/KyleBanks/goggles
WORKDIR $GOPATH/src/github.com/KyleBanks/goggles
ADD . .

CMD go test -cover $(go list ./... | grep -v vendor)