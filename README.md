# Gom (Parser Combinator Library)

A lightweight parser combinator library for Go, inspired by the design philosophy of Rust’s [Nom](https://github.com/geal/nom). This library provides simple building blocks (`Tag`, `Sequence`, `Alt`, etc.) that can be combined to parse structured data with minimal boilerplate. This aims to have every string parser that nom provides, at minimum.

---

## Installation

```bash
go get github.com/SamYaple/gom
```

---

## Usage

Please see [`example/main.go`](example/main.go) in this repository.

### Quick Example

```go
package main

import (
    "fmt"
    "github.com/SamYaple/gom"
)

func main() {
    input := "Hello World!"

    // Create a parser that matches "Hello", then " ", followed by either "Bob" or "World"
    parser := gom.Sequence(
        gom.Tag("Hello"),
        gom.Tag(" "),
        gom.Alt(
            gom.Tag("Bob"),
            gom.Tag("World"),
        ),
    )

    remaining, parsed, err := parser(input)
    if err != nil {
        fmt.Println("Parsing error:", err)
        return
    }
    fmt.Println("Parsed:", parsed)
    fmt.Println("Remaining: ", remaining)
}
```

Running this should print:
```
Parsed: [Hello  World]
Remaining: !
```
