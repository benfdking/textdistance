# textdistance

[![](https://godoc.org/github.com/benfdking/textdistance?status.svg)](http://godoc.org/github.com/benfdking/textdistance)
[![Go Report Card](https://goreportcard.com/badge/github.com/benfdking/textdistance)](https://goreportcard.com/report/github.com/benfdking/textdistance)

`textdistance` is a string comparison library written in Go. It is heavily `inspired` by the identically named Python library. There is a slight difference though that this library aims to be written in a `safe` way such that it does not panic.

## Usage

### Install

```bash
go get -u github.com/benfdking/textdistance
```

### Example

```go
package main 

import "github.com/benfdking/textdistance"

func main(){
	h := textdistance.Hamming{}
	distance := h.Distance("drummer", "dresser")
	fmt.Println(distance)
}
```


## Documentation

The full documentation with further examples is available at [GoDoc](https://godoc.org/github.com/benfdking/textdistance).
