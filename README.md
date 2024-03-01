# go-hexpalette

Package **hexpalette** provides tools for working with a palette specified with **hex color palette codes**, for the Go programming language.

**Hex Color Codes** look like these:

* `#1B2A34`
* `1E5AA8`
* `#fd0304`
* `00852b`

A **hex color palette code** looks like these:

* `e63946edae493376bd00798c52489c`
* `f9731b7e2a9c008dd3`
* `f4538afaa300f5dd6159d5e0`
* `4e3806f0a501eac016f074004e3809`
* `ea5545f46a9bef9b20edbf33ede15bbdcf3287bc4527aeefb33dc6`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-hexpalette

[![GoDoc](https://godoc.org/github.com/reiver/go-hexpalette?status.svg)](https://godoc.org/github.com/reiver/go-hexpalette)

## Example

Here is a simple example of parsing a hex color code:

```golang

var code string = "#ea5545 #F46a9B ef9b20    edbf33 #ede15b #BDCF32 #87bC45 27aeef #b33dc6  "

code, err := hexpalette.CanonicalString(code)
```

## Import

To import package **hexpalette** use `import` code like the follownig:
```
import "github.com/reiver/go-hexpalette"
```

## Installation

To install package **hexpalette** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-hexpalette
```

## Author

Package **hexpalette** was written by [Charles Iliya Krempeaux](http://changelog.ca)
