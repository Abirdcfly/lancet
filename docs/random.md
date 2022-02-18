# Random
Package random implements some basic functions to generate random int and string.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/main/random/random.go](https://github.com/duke-git/lancet/blob/main/random/random.go)


<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/random"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [RandBytes](#RandBytes)
- [RandInt](#RandInt)
- [RandString](#RandString)

<div STYLE="page-break-after: always;"></div>

## Documentation


### <span id="RandBytes">RandBytes</span>
<p>Generate random byte slice.</p>

<b>Signature:</b>

```go
func RandBytes(length int) []byte
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
	randBytes := random.RandBytes(4)
	fmt.Println(randBytes)
}
```


### <span id="RandInt">RandInt</span>
<p>Generate random int between min and max, may contain min, not max.</p>

<b>Signature:</b>

```go
func RandInt(min, max int) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
	rInt := random.RandInt(1, 10)
	fmt.Println(rInt)
}
```



### <span id="RandString">RandInt</span>
<p>Generate random given length string.</p>

<b>Signature:</b>

```go
func RandString(length int) string
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
	randStr := random.RandString(6)
	fmt.Println(randStr)
}
```


