# go-uniseg

Package `uniseg` implements basic unicode segmentation.

Documentation at <https://godoc.org/github.com/tim-st/go-uniseg>.

Download package `uniseg` with `go get -u github.com/tim-st/go-uniseg/...`

## Example

```go
package main

import (
	"fmt"

	"github.com/tim-st/go-uniseg"
)

func main() {
	const text = `Some test code. 10 Cafés, GitHub.`

	segments := uniseg.Segments([]byte(text))

	for idx, segment := range segments {
		fmt.Println(idx, segment)
	}
}

```

```
0 Segment{"Some" (Category: Wf; Runes: 4)}
1 Segment{" " (Category: Zs; Runes: 1)}
2 Segment{"test" (Category: Wl; Runes: 4)}
3 Segment{" " (Category: Zs; Runes: 1)}
4 Segment{"code" (Category: Wl; Runes: 4)}
5 Segment{"." (Category: Po; Runes: 1)}
6 Segment{" " (Category: Zs; Runes: 1)}
7 Segment{"10" (Category: Nd; Runes: 2)}
8 Segment{" " (Category: Zs; Runes: 1)}
9 Segment{"Cafés" (Category: Wf; Runes: 5)}
10 Segment{"," (Category: Po; Runes: 1)}
11 Segment{" " (Category: Zs; Runes: 1)}
12 Segment{"GitHub" (Category: Wm; Runes: 6)}
13 Segment{"." (Category: Po; Runes: 1)}
```