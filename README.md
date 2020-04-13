# textdistance

[![](https://godoc.org/github.com/benfdking/textdistance?status.svg)](http://godoc.org/github.com/benfdking/textdistance)
[![Go Report Card](https://goreportcard.com/badge/github.com/benfdking/textdistance)](https://goreportcard.com/report/github.com/benfdking/textdistance)

`textdistance` is a string comparison library written in Go. Heavily inspired by the identically named [Python library](https://github.com/life4/textdistance), it aims to provide a myriad of different algorithms. 
 
Additionally, it aims to be:

- safe for production use, preferring `error` where required over absolute raw performance
- have consistent interfaces in order to test different implementations to enable easy dependency injection
- within those constraints being as performant as possible: allowing assembly snippets but no `C` library links 

## Documentation

The full documentation with further examples is available at [GoDoc](https://godoc.org/github.com/benfdking/textdistance).

## Usage

### Install

```bash
go get -u github.com/benfdking/textdistance
```

### Example

```go
package main 

import (
    "fmt"

    "github.com/benfdking/textdistance"
)

func main(){
	h := textdistance.Hamming{}
	distance, _ := h.Distance("drummer", "dresser")
	fmt.Println(distance)
}
```
