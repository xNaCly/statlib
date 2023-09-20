## statlib

Go library for commonly used statistical computations, optimized for performance

## Usage

Computing the value of 12!:

```go
package main

import (
   "fmt"
   "github.com/xnacly/statlib"
)

func main() {
    val, err := statlib.Fac(12)
    if err != nil {
        panic(err)
    }
    fmt.Printf("12! = %d\n", val)
}
```

Choosing 6 possibilities out of 49 options:

```go
package main

import (
   "fmt"
   "github.com/xnacly/statlib"
)

func main() {
    val, err := statlib.BinomialCoefficient(49, 6)
    if err != nil {
        panic(err)
    }
    fmt.Printf("49nCr6 = %d\n", val)
}
```

Calculating the chance of getting 3 heads when throwing a coin 10 times:

```go
package main

import (
   "fmt"
   "github.com/xnacly/statlib/distributions"
)

func main() {
	bin, err := BinomialNew(10, 0.5)
	if err != nil {
		panic(err)
	}

	val, err := bin.Prob(3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("chance of getting 3 heads when throwing 10 times = %f%%\n", val*100)
}
```

## Features

- [x] Factorial
- [x] Binomial coefficient
- [ ] Distributions (for `X=k`)
  - [ ] for each of the following
    - [ ] Median
    - [ ] Mode
    - [ ] Mean
    - [ ] Variance
  - [x] Binomial
  - [ ] Normal
  - [ ] Poisson

## Contributing

Make sure to write understandable code, document hacks, tricks and hard to read
sections with the source and or an explanation. Always provide tests, focus
especially on edge cases and large inputs.

Consider well written commits:

```
<file without extension or context>: <short description about the change>

<
longer description, go into depth on your implementation and mention github issues with:

Fix: #00000 if fixed
Add: #00000 if added a feature
Update: #00000 if changes regarding the issue occured in your commit
>
```

## Test

Run the following to execute the test suite

```sh
go test ./...
```
