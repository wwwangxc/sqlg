[![Security Status](https://www.murphysec.com/platform3/v3/badge/1609921498488274944.svg)](https://www.murphysec.com/accept?code=299ad9b983a705c0e56686a746816921&type=1&from=2)

[![Go Report Card](https://goreportcard.com/badge/github.com/wwwangxc/sqlg)](https://goreportcard.com/report/github.com/wwwangxc/sqlg)
[![GoDoc](https://pkg.go.dev/badge/github.com/wwwangxc/sqlg?status.svg)](https://pkg.go.dev/github.com/wwwangxc/sqlg)

# SQLg

An easy way to generate SQL statements for Go. 🤗

## Install

```sh
go get github.com/wwwangxc/sqlg
```

## Quick Start

### Select

```go
package main

import (
        "fmt"

        "github.com/com/wwwangxc/sqlg"
)

func main () {
        // SELECT * FROM user WHERE id=? AND deleted_at IS NULL
        // [666]
        g := sqlg.NewGenerator("user",
                sqlg.WithAnd("id", sqlg.EQ(666)),
                sqlg.WithAnd("deleted_at", sqlg.Null()))
        _, _ = g.Select()
    
        // SELECT * FROM user WHERE id!=? AND deleted_at IS NOT NULL
        // [666]
        g = sqlg.NewGenerator("user",
                sqlg.WithAnd("id", sqlg.EQ(666)),
                sqlg.WithAnd("deleted_at", sqlg.Null()))
        _, _ = g.Select()
    
        // SELECT id, name FROM user WHERE id>=? OR name=? ORDER BY id DESC LIMIT 10
        // [666 tom]
        g = sqlg.NewGenerator("user",
                sqlg.WithAnd("id", sqlg.GTE(666)),
                sqlg.WithOr("name", sqlg.EQ("tom")),
                sqlg.WithOrderByDESC("id"), sqlg.WithLimit(10))
        _, _ = g.Select("id", "name")
    
        type User struct {
                ID   uint64 `sqlg:"id"`
                Name string `sqlg:"name"`
        }
    
        // compound expression
        m := sqlg.NewCompExpr()
        m.Put("name", sqlg.EQ("tom"))
        m.Put("id", sqlg.EQ(666))
        
        // SELECT id, name FROM user WHERE deleted_at IS NULL AND (name=? OR id=?)
        // [tom 666]
        // <nil>
        g = sqlg.NewGenerator("user",
                sqlg.WithAnd("deleted_at", sqlg.Null()),
                sqlg.WithAndExprs(m))
        _, _, _ = g.SelectByStruct(User{})
}
```
