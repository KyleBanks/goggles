#!/bin/bash
set -e

cd $GOPATH/src/github.com/KyleBanks/go-kit

go get -u github.com/golang/lint/golint
go get -u github.com/robertkrimen/godocdown/godocdown

echo "-------------  TEST  -------------"
go test -cover $@ $(go list ./... | grep -v vendor/)

echo "-------------  VET  -------------"
go vet $(go list ./... | grep -v vendor/)

echo "-------------  LINT  -------------"
golint $(go list ./... | grep -v vendor/)

echo "-------------  FMT  -------------"
go fmt $(go list ./... | grep -v vendor/)

echo "-------------  DOCS  -------------"
doc_file="README.md"

# Write the README header
cat <<EOF > $doc_file
# go-kit

[![Build Status](https://travis-ci.org/KyleBanks/go-kit.svg?branch=master)](https://travis-ci.org/KyleBanks/go-kit) &nbsp;
[![GoDoc](https://godoc.org/github.com/KyleBanks/go-kit?status.svg)](https://godoc.org/github.com/KyleBanks/go-kit) &nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/go-kit)](https://goreportcard.com/report/github.com/KyleBanks/go-kit) &nbsp;
[![Coverage Status](https://coveralls.io/repos/github/KyleBanks/go-kit/badge.svg?branch=master)](https://coveralls.io/github/KyleBanks/go-kit?branch=master)

This repository contains generic Go packages that are reused throughout various Go projects. 

## Packages

Most packages are designed to be used standalone, however a few such as \`auth\` have additional dependencies on other packages in the \`go-kit\`. 

EOF

# For each package, write the summary
for d in */ ; do
	if [ "$d" == "vendor/" ]; then
		continue
	fi
    
    godocdown -template=package.tmpl "github.com/KyleBanks/go-kit/$d" >> "$doc_file"
    echo "" >> "$doc_file"

    godocdown -no-template "github.com/KyleBanks/go-kit/$d" > "./$d/$doc_file"
done

# Finally, add the README footer
cat <<EOF >> "$doc_file"
## Testing

\`\`\`
./sanity.sh
\`\`\`

## License

\`\`\`
$(cat LICENSE)
\`\`\`
EOF