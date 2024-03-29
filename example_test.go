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
		ID     uint64 `db:"id"`
		Name   string `db:"name"`
		Age    uint8
		Height uint8 `db:"-"`
	}
	compExpr := sqlg.NewCompExpr()
	compExpr.Put("name", sqlg.EQ("tom"))
	compExpr.Put("id", sqlg.EQ(666))
	g = sqlg.NewGenerator("user", sqlg.WithAnd("deleted_at", sqlg.Null()), sqlg.WithAndExprs(compExpr))
	sql, params, err := g.SelectByStruct(User{})
	fmt.Println("*Return-4")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Return-5
	g = sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.EQ(666)), sqlg.WithLimit(1))
	assExpr := sqlg.NewAssExpr()
	assExpr.Put("name", "jerry")
	assExpr.Put("age", 3)
	sql, params = g.Update(assExpr)
	fmt.Println("*Return-5")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-6
	g = sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.EQ(666)), sqlg.WithLimit(1))
	sql, params = g.Delete()
	fmt.Println("*Return-6")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-7
	g = sqlg.NewGenerator("user")
	columns := []string{"name", "age"}
	records := [][]interface{}{
		{"tom", 5},
		{"jerry", 3},
	}
	sql, params = g.Insert(columns, records...)
	fmt.Println("*Return-7")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-8
	assExpr = sqlg.NewAssExpr()
	assExpr.Put("name", "tom")
	assExpr.Put("age", "5")
	g = sqlg.NewGenerator("user", sqlg.OnDuplicateKeyUpdate(assExpr))
	columns = []string{"name", "age"}
	records = [][]interface{}{{"tom", 5}}
	sql, params = g.Insert(columns, records...)
	fmt.Println("*Return-8")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-9
	compExpr = sqlg.NewCompExpr()
	compExpr.Put("name", sqlg.EQ("tom"))
	compExpr.Put("age", sqlg.GTE(5))
	g = sqlg.NewGenerator("user", sqlg.WithNExists("user", compExpr))
	columns = []string{"name", "age"}
	records = [][]interface{}{{"tom", 5}}
	sql, params = g.Insert(columns, records...)
	fmt.Println("*Return-9")
	fmt.Println(sql)
	fmt.Println(params)

	// Output:
	// *Return-1
	// SELECT * FROM `user` WHERE `id`=? AND `deleted_at` IS NULL
	// [666]
	//
	// *Return-2
	// SELECT * FROM `user` WHERE `id`!=? AND `deleted_at` IS NOT NULL
	// [666]
	//
	// *Return-3
	// SELECT `id`, `name` FROM `user` WHERE `id`>=? OR `name`=? ORDER BY `id` DESC LIMIT 10
	// [666 tom]
	//
	// *Return-4
	// SELECT `id`, `name` FROM `user` WHERE `deleted_at` IS NULL AND (`name`=? OR `id`=?)
	// [tom 666]
	// <nil>
	//
	// *Return-5
	// UPDATE `user` SET `name`=?, `age`=? WHERE `id`=? LIMIT 1
	// [jerry 3 666]
	//
	// *Return-6
	// DELETE FROM `user` WHERE `id`=? LIMIT 1
	// [666]
	//
	// *Return-7
	// INSERT INTO `user` (`name`, `age`) VALUES (?,?), (?,?)
	// [tom 5 jerry 3]
	//
	// *Return-8
	// INSERT INTO `user` (`name`, `age`) VALUES (?,?) ON DUPLICATE KEY UPDATE `name`=?, `age`=?
	// [tom 5 tom 5]
	//
	// *Return-9
	// INSERT INTO `user` (`name`, `age`) SELECT ?,? FROM dual WHERE NOT EXISTS (SELECT * FROM `user` WHERE `name`=? AND `age`>=?)
	// [tom 5 tom 5]
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

	// generate SELECT sql statement and params
	sql, params = g.Select("id", "name")
	fmt.Println("*Return-4")
	fmt.Println(sql)
	fmt.Println(params)

	// Output:
	// *Return-1
	// SELECT * FROM `user` WHERE `id`=? AND `deleted_at` IS NULL
	// [666]
	//
	// *Return-2
	// SELECT * FROM `user` WHERE `id`!=? AND `deleted_at` IS NOT NULL
	// [666]
	//
	// *Return-3
	// SELECT `id`, `name` FROM `user` WHERE `id`>=? OR `name`=? ORDER BY `id` DESC LIMIT 10
	// [666 tom]
	//
	// *Return-4
	// SELECT `id`, `name` FROM `user` WHERE `deleted_at` IS NULL AND (`name`=? OR `id`=?)
	// [tom 666]
}

func ExampleGenerator_SelectByStruct() {
	type User struct {
		ID     uint64 `db:"id"`
		Name   string `db:"name"`
		Age    uint8
		Height uint8 `db:"-"`
	}

	// compound expression
	m := sqlg.NewCompExpr()
	m.Put("name", sqlg.EQ("tom"))
	m.Put("id", sqlg.EQ(666))

	// create generator
	g := sqlg.NewGenerator("user", sqlg.WithAnd("deleted_at", sqlg.Null()), sqlg.WithAndExprs(m))

	// Return-1
	// generate SELECT sql statement and params
	sql, params, err := g.SelectByStruct(nil)
	fmt.Println("*Return-1")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Return-2
	// generate SELECT sql statement and params
	sql, params, err = g.SelectByStruct(User{})
	fmt.Println("*Return-2")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Return-3
	// generate SELECT sql statement and params
	sql, params, err = g.SelectByStruct(&User{})
	fmt.Println("*Return-3")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()

	// Return-4
	// generate SELECT sql statement and params
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
	// SELECT `id`, `name` FROM `user` WHERE `deleted_at` IS NULL AND (`name`=? OR `id`=?)
	// [tom 666]
	// <nil>
	//
	// *Return-3
	// SELECT `id`, `name` FROM `user` WHERE `deleted_at` IS NULL AND (`name`=? OR `id`=?)
	// [tom 666]
	// <nil>
	//
	// *Return-4
	// SELECT `id`, `name` FROM `user` WHERE `deleted_at` IS NULL AND (`name`=? OR `id`=?)
	// [tom 666]
	// <nil>
}

func ExampleGenerator_Update() {
	// create generator
	g := sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.EQ(666)), sqlg.WithLimit(1))

	// assignment expression
	assExpr := sqlg.NewAssExpr()
	assExpr.Put("name", "jerry")
	assExpr.Put("age", 3)

	// generate UPDATE sql statement and params
	sql, params := g.Update(assExpr)
	fmt.Println(sql)
	fmt.Println(params)

	// Output:
	// UPDATE `user` SET `name`=?, `age`=? WHERE `id`=? LIMIT 1
	// [jerry 3 666]
}

func ExampleGenerator_Delete() {
	// create generator
	g := sqlg.NewGenerator("user", sqlg.WithAnd("id", sqlg.EQ(666)), sqlg.WithLimit(1))

	// generate DELETE sql statement and params
	sql, params := g.Delete()
	fmt.Println(sql)
	fmt.Println(params)

	// Output:
	// DELETE FROM `user` WHERE `id`=? LIMIT 1
	// [666]
}

func ExampleGenerator_Insert() {
	// Return-1
	// create generator
	g := sqlg.NewGenerator("user")
	columns := []string{"name", "age"}
	records := [][]interface{}{
		{"tom", 5},
		{"jerry", 3},
	}

	// generate INSERT sql statment and params
	// INSERT INTO ...
	sql, params := g.Insert(columns, records...)
	fmt.Println("*Return-1")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-2
	// create assignment expression
	assExpr := sqlg.NewAssExpr()
	assExpr.Put("name", "tom")
	assExpr.Put("age", "5")

	// create generator
	g = sqlg.NewGenerator("user", sqlg.OnDuplicateKeyUpdate(assExpr))
	columns = []string{"name", "age"}
	records = [][]interface{}{{"tom", 5}}

	// generate INSERT sql statment and params
	// INSERT INTO ... ON DUPLICATE KEY UPDATE ...
	sql, params = g.Insert(columns, records...)
	fmt.Println("*Return-2")
	fmt.Println(sql)
	fmt.Println(params)
	fmt.Println()

	// Return-3
	// create compound expression
	compExpr := sqlg.NewCompExpr()
	compExpr.Put("name", sqlg.EQ("tom"))
	compExpr.Put("age", sqlg.GTE(5))

	// create generator
	g = sqlg.NewGenerator("user", sqlg.WithNExists("user", compExpr))
	columns = []string{"name", "age"}
	records = [][]interface{}{{"tom", 5}}

	// generate INSERT sql statment and params
	// INSERT INTO ... WHERE NOT EXIST (...)
	sql, params = g.Insert(columns, records...)
	fmt.Println("*Return-3")
	fmt.Println(sql)
	fmt.Println(params)

	// Output:
	// *Return-1
	// INSERT INTO `user` (`name`, `age`) VALUES (?,?), (?,?)
	// [tom 5 jerry 3]
	//
	// *Return-2
	// INSERT INTO `user` (`name`, `age`) VALUES (?,?) ON DUPLICATE KEY UPDATE `name`=?, `age`=?
	// [tom 5 tom 5]
	//
	// *Return-3
	// INSERT INTO `user` (`name`, `age`) SELECT ?,? FROM dual WHERE NOT EXISTS (SELECT * FROM `user` WHERE `name`=? AND `age`>=?)
	// [tom 5 tom 5]

}
