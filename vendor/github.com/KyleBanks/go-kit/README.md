# go-kit

[![Build Status](https://travis-ci.org/KyleBanks/go-kit.svg?branch=master)](https://travis-ci.org/KyleBanks/go-kit) &nbsp;
[![GoDoc](https://godoc.org/github.com/KyleBanks/go-kit?status.svg)](https://godoc.org/github.com/KyleBanks/go-kit) &nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/go-kit)](https://goreportcard.com/report/github.com/KyleBanks/go-kit) &nbsp;
[![Coverage Status](https://coveralls.io/repos/github/KyleBanks/go-kit/badge.svg?branch=master)](https://coveralls.io/github/KyleBanks/go-kit?branch=master)

This repository contains generic Go packages that are reused throughout various Go projects. 

## Packages

Most packages are designed to be used standalone, however a few such as `auth` have additional dependencies on other packages in the `go-kit`. 

### [auth](./auth)

`import "github.com/KyleBanks/go-kit/auth/"`

Package auth provides generic authentication functionality.

Note this is not 100% secure and should only be used for prototyping, not for
production systems or systems that are accessed by real users.

### [cache](./cache)

`import "github.com/KyleBanks/go-kit/cache/"`

Package cache is a simple cache wrapper, used to abstract Redis/Memcache/etc
behind a reusable API for simple use cases.

The idea is that Redis could be swapped for another cache and the client
wouldn't need to update another (except perhaps calls to New to provide
different connection parameters).

For now cache supports only Redis, but eventually that could be provided by the
client.

### [clipboard](./clipboard)

`import "github.com/KyleBanks/go-kit/clipboard/"`

Package clipboard provides the ability to read and write to the system
clipboard.

Note: Currently only supports Mac OS.

### [contains](./contains)

`import "github.com/KyleBanks/go-kit/contains/"`

Package contains adds a few small helper functions to see if a slice contains a
particular value.

### [convert](./convert)

`import "github.com/KyleBanks/go-kit/convert/"`

Package convert provides generalized type conversion utilities.

### [env](./env)

`import "github.com/KyleBanks/go-kit/env/"`

Package env provides application environment detection, and support for a
Dev/Test/Prod environment system.

### [git](./git)

`import "github.com/KyleBanks/go-kit/git/"`

Package git provides git source control functionality.

### [gonamo](./gonamo)

`import "github.com/KyleBanks/go-kit/gonamo/"`

Package gonamo provides a simple wrapper around the DynamoDB SDK.

The intention is to provide a minimal DynamoDB table representation that can be
created, written to, queried and scanned.

The main starting point of a gonamo implementation is to define a model that you
want to store in DynamoDB, and implement the "Persistable" interface:

    type Repository struct {
    	Owner string
    	Name  string

    	DateCreated time.Time
    }

    func (r Repository) HashKey() interface{} {
    	return r.Owner
    }

    func (r Repository) RangeKey() interface{} {
    	return r.Name
    }

    func (r Repository) Attributes() gonamo.AttributeMap {
    	return := gonamo.AttributeMap{
    		"owner":       gonamo.AttributeValue(gonamo.StringType, r.Owner),
    		"name":        gonamo.AttributeValue(gonamo.StringType, r.Name),
    		"dateCreated": gonamo.AttributeValue(gonamo.NumberType, r.DateCreated.Unix()),
    	}
    }

Next you will be able to create a Table using the NewTable method, by providing
a table name and the key structure:

    gonamo.HashRangeKeyDefinition{"owner", gonamo.StringType, "name", gonamo.StringType}
    tbl, err := gonamo.NewTable("tableName", key, nil)

### [job](./job)

`import "github.com/KyleBanks/go-kit/job/"`

Package job provides the ability to execute tasks on a timed interval.

### [log](./log)

`import "github.com/KyleBanks/go-kit/log/"`

Package log provides a simple logging service to print to stdout/stderr with
timestamp and log source information.

### [milliseconds](./milliseconds)

`import "github.com/KyleBanks/go-kit/milliseconds/"`

### [orm](./orm)

`import "github.com/KyleBanks/go-kit/orm/"`

Package orm manages access to a database, including ORM-like functionality.

The package wraps the GORM library, which can then be potentially swapped out
with minimal changes.

### [push](./push)

`import "github.com/KyleBanks/go-kit/push/"`

Package push provides GCM and APN push notification functionality.

### [router](./router)

`import "github.com/KyleBanks/go-kit/router/"`

Package router defines the Route interface, and registers routes to an http
server.

### [storage](./storage)

`import "github.com/KyleBanks/go-kit/storage/"`

Package storage provides the ability to persist and retrieve structs.

### [timer](./timer)

`import "github.com/KyleBanks/go-kit/timer/"`

Package timer provides the ability to time abritrary events, like the duration
of a method call.

### [today](./today)

`import "github.com/KyleBanks/go-kit/today/"`

Package today provides utilities for access regarding today's date.

### [unique](./unique)

`import "github.com/KyleBanks/go-kit/unique/"`

Package unique provides the functionality to create unique versions of slices.

## Testing

```
./sanity.sh
```

## License

```
The MIT License (MIT)

Copyright (c) 2017 Kyle Banks

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
