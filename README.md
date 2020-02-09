# textdistance

[![](https://godoc.org/github.com/benfdking/textdistance?status.svg)](http://godoc.org/github.com/benfdking/textdistance)

`textdistance` is a string comparison library written in Go

## Install

```bash
go get -u github.com/benfdking/textdistance
```

## Example

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
