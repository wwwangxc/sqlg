package sqlg_test

import (
	"fmt"

	"github.com/wwwangxc/sqlg"
)

func Example() {
	// Return-1
	g := sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.EQ(666)), sqlg.WithAnd("deleted_at", sqlg.Null()))
	sql, params := g.Select()
	fmt.Println("*Return-1")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-2
	g = sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.NEQ(666)), sqlg.WithAnd("deleted_at", sqlg.NNull()))
	sql, params = g.Select()
	fmt.Println("*Return-2")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-3
	g = sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.GTE(666)), sqlg.WithOr("name", sqlg.EQ("tom")), sqlg.WithOrderByDESC("id"), sqlg.WithLimit(10))
	sql, params = g.Select("id", "name")
	fmt.Println("*Return-3")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-4
	type User struct {
		ID   uint64 `sqlg:"id"`
		Name string `sqlg:"name"`
	}
	m := sqlg.NewCompExpr()
	m.Put("name", sqlg.EQ("tom"))
	m.Put("id", sqlg.EQ(666))
	g = sqlg.NewGenerator("user", sqlg.WithAnd("deleted_at", sqlg.Null()), sqlg.WithAndExprs(m))
	sql, params, err := g.SelectByStruct(User{})
	fmt.Println("*Return-4")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)

	// Output:
	// *Return-1
	// SELECT * FROM user WHERE id=? AND deleted_at IS NULL
	// [666]
	//
	// *Return-2
	// SELECT * FROM user WHERE id!=? AND deleted_at IS NOT NULL
	// [666]
	//
	// *Return-3
	// SELECT id, name FROM user WHERE id>=? OR name=? ORDER BY id DESC LIMIT 10
	// [666 tom]
	//
	// *Return-4
	// SELECT id, name FROM user WHERE deleted_at IS NULL AND (name=? OR id=?)
	// [tom 666]
	// <nil>
}

func ExampleGenerator_Select() {
	// Return-1
	g := sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.EQ(666)), sqlg.WithAnd("deleted_at", sqlg.Null()))
	sql, params := g.Select()
	fmt.Println("*Return-1")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-2
	g = sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.NEQ(666)), sqlg.WithAnd("deleted_at", sqlg.NNull()))
	sql, params = g.Select()
	fmt.Println("*Return-2")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-3
	g = sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.GTE(666)), sqlg.WithOr("name", sqlg.EQ("tom")), sqlg.WithOrderByDESC("id"), sqlg.WithLimit(10))
	sql, params = g.Select("id", "name")
	fmt.Println("*Return-3")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-4
	// compound expression
	m := sqlg.NewCompExpr()
	m.Put("name", sqlg.EQ("tom"))
	m.Put("id", sqlg.EQ(666))

	// create generator
	g = sqlg.NewGenerator("user", sqlg.WithAnd("deleted_at", sqlg.Null()), sqlg.WithAndExprs(m))

	// generate sql statement and params
	sql, params = g.Select("id", "name")
	fmt.Println("*Return-4")
	fmt.Println(sql)
	fmt.Println(params)

	// Output:
	// *Return-1
	// SELECT * FROM user WHERE id=? AND deleted_at IS NULL
	// [666]
	//
	// *Return-2
	// SELECT * FROM user WHERE id!=? AND deleted_at IS NOT NULL
	// [666]
	//
	// *Return-3
	// SELECT id, name FROM user WHERE id>=? OR name=? ORDER BY id DESC LIMIT 10
	// [666 tom]
	//
	// *Return-4
	// SELECT id, name FROM user WHERE deleted_at IS NULL AND (name=? OR id=?)
	// [tom 666]
}

func ExampleGenerator_SelectByStruct() {
	type User struct {
		ID   uint64 `sqlg:"id"`
		Name string `sqlg:"name"`
	}

	// compound expression
	m := sqlg.NewCompExpr()
	m.Put("name", sqlg.EQ("tom"))
	m.Put("id", sqlg.EQ(666))

	// create generator
	g := sqlg.NewGenerator("user", sqlg.WithAnd("deleted_at", sqlg.Null()), sqlg.WithAndExprs(m))

	// Return-1
	// generate sql statement and params
	sql, params, err := g.SelectByStruct(nil)
	fmt.Println("*Return-1")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Return-2
	// generate sql statement and params
	sql, params, err = g.SelectByStruct(User{})
	fmt.Println("*Return-2")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Return-3
	// generate sql statement and params
	sql, params, err = g.SelectByStruct(&User{})
	fmt.Println("*Return-3")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Return-4
	// generate sql statement and params
	var user *User
	sql, params, err = g.SelectByStruct(user)
	fmt.Println("*Return-4")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Output:
	// *Return-1
	//
	// []
	// target can not be empty
	//
	// *Return-2
	// SELECT id, name FROM user WHERE deleted_at IS NULL AND (name=? OR id=?)
	// [tom 666]
	// <nil>
	//
	// *Return-3
	// SELECT id, name FROM user WHERE deleted_at IS NULL AND (name=? OR id=?)
	// [tom 666]
	// <nil>
	//
	// *Return-4
	// SELECT id, name FROM user WHERE deleted_at IS NULL AND (name=? OR id=?)
	// [tom 666]
	// <nil>
}
