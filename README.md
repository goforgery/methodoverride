# methodoverride

[![Build Status](https://secure.travis-ci.org/goforgery/methodoverride.png?branch=master)](http://travis-ci.org/goforgery/methodoverride)

Method override service for [Forgery2](https://github.com/goforgery/forgery2).

## Install

	go get github.com/goforgery/methodoverride

## Use

Changes the HTTP request method to that set by the header `X-HTTP-Method-Override`.

```javascript
package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/methodoverride"
)

func main() {
	app := f.CreateApp()
	app.Use(methodoverride.Create())
	app.Listen(3000)
}
```

## Test

    go test
