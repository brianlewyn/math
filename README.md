# brianlewyn/math

The `brianlewyn/math` package is intended for programmers working with mathematical functions in the Go language. That is why this package provides functions that allow us to calculate arithmetic type functions, including derivatives and integrals. The goal is that from a text string containing a mathematical expression, the result can be returned in the same way.

---
* [Install](#install)
* [Package Properties](#package-properties)
* [Arithmetic](#arithmetic)
* [Derivatives](#derivatives)
* [Integrals](#integrals)
* [License](#license)
---



## Install
With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/brianlewyn/math
```

## Package Properties
### Literal: `x`
* It's case sensitive.
* Only one letter is accepted.
* It can be any letter of the American alphabet.

### Expression: `gx`
* If you enter `+kx^+n` as `kxn`, the program you can infer and reconstruct the expressions to their base form, to do the operations correctly.
* For both the constant and the exponent, you can use integers or floats.
* Negative signs are taken into account, both for constants and for the exponent.
* Between each positive expression you can put `x x`, `x +x` or `x + x`.
* Between each negative expression you can put `x -x` or `x - x`.

## Arithmetic
### Arithmetic.Add()
* As long as the same literal and exponent exist, only the expressions that meet the condition will be added.
* Sort the expressions from highest to lowest exponent.
```go
package main

import "fmt"
import "github.com/brianlewyn/math/arithmetic"

func main() {
	x, gx := "x", "x^8.5 x^2 .5 9x^8.5 9x^2 9.5"
	err := arithmetic.Add(&x, &gx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("f(%s) = %s\n", x, gx)
	}
}
```
```bash
Output
f(x) = 10x^8.5 +10x^2 +10x^0
```

### Arithmetic.Multiply()
* You must always use parentheses, to be able to multiply.
* Can only multiply with a single hierarchy level.
```go
package main

import "fmt"
import "github.com/brianlewyn/math/arithmetic"

func main() {
	x, gx := "x", "(.5x)(5x^7.5 10x 20x-1)"
	err := arithmetic.Multiply(&x, &gx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("f(%s) = %s\n", x, gx)
	}
}
```
```bash
Output
f(x) = 2.5x^8.5 +5x^2 +10x^0
```

## Derivatives
### Derivatives.Add()
* Only simple expressions are accepted.
```go
package main

import "fmt"
import "github.com/brianlewyn/math/derivatives"

func main() {
	x, gx := "x", "3x9 3x3 3x 33"
	err := derivatives.Add(&x, &gx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("f(%s) = %s\n", x, gx)
	}
}
```
```bash
Output
f(x) = 27x^8 +9x^2 +3x^0
```

## Integrals
Still in development

## License
MIT License. See the LICENSE file for details.
