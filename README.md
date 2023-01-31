# Go SQL Like expression
Convert SQL Like patterns to Go regular expressions.

[![Test](https://github.com/shellyln/go-sql-like-expr/actions/workflows/test.yml/badge.svg)](https://github.com/shellyln/go-sql-like-expr/actions/workflows/test.yml)
[![release](https://img.shields.io/github/v/release/shellyln/go-sql-like-expr)](https://github.com/shellyln/go-sql-like-expr/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/shellyln/go-sql-like-expr)](https://github.com/shellyln/go-sql-like-expr)

## 🚀 Usage

```go
package main

import (
    "fmt"
    "regexp"
    "github.com/shellyln/go-sql-like-expr/likeexpr"
)

func main() {
    // Parameters:
    // Like pattern, Escape character, Escaping the escape character itself
    pat := likeexpr.ToRegexp("fo%ba_", '\\', false)

    re, err := regexp.Compile(pat)
    if err != nil {
        fmt.Printf("Compile: error = %v\n", err)
        return
    }

    matched := re.MatchString("fooobar")
    fmt.Printf("Compile: matched = %v\n", matched)
}
```

## ⚖️ License

MIT  
Copyright (c) 2023 Shellyl_N and Authors.
