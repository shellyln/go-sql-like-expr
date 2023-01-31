# Go SQL Like expression
Convert SQL Like patterns to Go regular expressions.

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
