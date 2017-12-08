# Calcgo

[![Build Status](https://travis-ci.org/relnod/calcgo.svg?branch=master)](https://travis-ci.org/relnod/calcgo)
[![codecov](https://codecov.io/gh/relnod/calcgo/branch/master/graph/badge.svg)](https://codecov.io/gh/relnod/calcgo)
[![Godoc](https://godoc.org/github.com/relnod/calcgo?status.svg)](https://godoc.org/github.com/relnod/calcgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/relnod/calcgo)](https://goreportcard.com/report/github.com/relnod/calcgo)

This is an experimental learning project, to better understand the process of
lexing and parsing.

## Description
Calcgo exposes a lexer, parser and interpreter to get tokens, an ast and the
result of a basic mathematical calculation. All three functions accept a
language L(G) defined [here](#grammar).

The calculations follow basic math rules, like "point before line" rule. To
break this rule it is possible to use brackets.
There needs to be at least one whitespace character between an operator an a
number. All other whitespace character get ignored by the lexer.

#### Lexer:
``` go
calcgo.Lex("(1 + 2) * 3")
```
#### Parser:
``` go
calcgo.Parse("(1 + 2) * 3")
```
#### Interpreter:
``` go
calcgo.Interpret("1 + 2 * 3")   // Result: 7
calcgo.Interpret("(1 + 2) * 3") // Result: 9
```

## Example
``` go
package main

import (
	"fmt"

	"github.com/relnod/calcgo"
)

func main() {
	number, _ := calcgo.Interpret("1 + 1")

	fmt.Println(number)
}
```

## Grammar

### Lexer:
The lexer accepts the language L(G), where G is the following deterministic
contextfree grammar:

G = (N,T,P,S)

N = {S}

T = {n, o, l, r}

P contains the following rules:

S → SOS

S → lSr

S →  n

S →  nN

O →  sos

N →  dF

N →  nF

N →  n

F →  nF

F →  n

#### Terminals
n ∈ { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' }

d ∈  { '.' }

o ∈ { '+', '-', '*'*, '/' }

l ∈ { '(' }

r ∈ { ')' }

s ∈  { ' ' }

## Tests and Benchmarks

#### Running Tests

Run tests with ```go test -v -race``` or with ```goconvey``` to see live test
results in a browser.

#### Running Benchmarks

Benchmarks can be tun with ```go test -run=^$ -bench=.```

To see the differences between current branch and master run ```./scripts/benchcmp.sh -n 5```

## License

This project is licensed under the MIT License. See the
[LICENSE](../master/LICENSE) file for details.
