package sqlg

import (
	"errors"
	"testing"
)

func TestGenerator_Select(t *testing.T) {
	g := NewGenerator("table_name")
	gotSQL, gotParams := g.Select("col1", "col2", "col3")

	wantSQL := "SELECT `col1`, `col2`, `col3` FROM `table_name`"
	wantParams := []interface{}{}

	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)

	m := NewCompExpr()
	m.Put("comp_col_eq", EQ("comp_val_eq"))
	m.Put("comp_col_neq", NEQ("comp_val_neq"))
	m.Put("comp_col_gt", GT("comp_val_gt"))
	m.Put("comp_col_gte", GTE("comp_val_gte"))
	m.Put("comp_col_lt", LT("comp_val_lt"))
	m.Put("comp_col_lte", LTE("comp_val_lte"))
	m.Put("comp_col_between", Between("comp_val_between1", "comp_val_between2"))
	m.Put("comp_col_not_between", NBetween("comp_val_not_between1", "comp_val_not_between2"))
	m.Put("comp_col_in", In([]interface{}{"comp_val_in1", "comp_val_in2", "comp_val_in3"}))
	m.Put("comp_col_not_in", NIn([]interface{}{"comp_val_not_in1", "comp_val_not_in2", "comp_val_not_in3"}))
	m.Put("comp_col_like", Like("comp_val_like"))
	m.Put("comp_col_not_like", NLike("comp_val_not_like"))
	m.Put("comp_col_like_prefix", LikePrefix("comp_val_like_prefix"))
	m.Put("comp_col_not_like_prefix", NLikePrefix("comp_val_not_like_prefix"))
	m.Put("comp_col_like_suffix", LikeSuffix("comp_val_like_suffix"))
	m.Put("comp_col_not_like_suffix", NLikeSuffix("comp_val_not_like_suffix"))
	m.Put("comp_col_null", Null())
	m.Put("comp_col_not_null", NNull())

	ops := []Option{
		WithAnd("col_eq", EQ("val_eq")),
		WithAnd("col_neq", NEQ("val_neq")),
		WithAnd("col_gt", GT("val_gt")),
		WithAnd("col_gte", GTE("val_gte")),
		WithAnd("col_lt", LT("val_lt")),
		WithAnd("col_lte", LTE("val_lte")),
		WithAnd("col_between", Between("val_between1", "val_between2")),
		WithAnd("col_not_between", NBetween("val_not_between1", "val_not_between2")),
		WithAnd("col_in", In([]interface{}{"val_in1", "val_in2", "val_in3"})),
		WithAnd("col_not_in", NIn([]interface{}{"val_not_in1", "val_not_in2", "val_not_in3"})),
		WithAnd("col_like", Like("val_like")),
		WithAnd("col_not_like", NLike("val_not_like")),
		WithAnd("col_like_prefix", LikePrefix("val_like_prefix")),
		WithAnd("col_not_like_prefix", NLikePrefix("val_not_like_prefix")),
		WithAnd("col_like_suffix", LikeSuffix("val_like_suffix")),
		WithAnd("col_not_like_suffix", NLikeSuffix("val_not_like_suffix")),
		WithAnd("col_null", Null()),
		WithAnd("col_not_null", NNull()),
		WithAndExprs(m),
		WithGroupBy("col_group_by_1", "col_group_by_2"),
		WithOrderBy("col_order_by"),
		WithOrderByDESC("col_order_by_desc"),
		WithLimit(666),
		WithOffset(999),
		ForceIndex("idx_some_index"),
		ForUpdate(),
	}

	g = NewGenerator("table_name", ops...)
	gotSQL, gotParams = g.Select("col1", "col2", "col3")

	wantSQL = "SELECT `col1`, `col2`, `col3` FROM `table_name` FORCE INDEX (`idx_some_index`) " +
		"WHERE `col_eq`=? " +
		"AND `col_neq`!=? " +
		"AND `col_gt`>? " +
		"AND `col_gte`>=? " +
		"AND `col_lt`<? " +
		"AND `col_lte`<=? " +
		"AND `col_between` BETWEEN ? AND ? " +
		"AND `col_not_between` NOT BETWEEN ? AND ? " +
		"AND `col_in` IN (?,?,?) " +
		"AND `col_not_in` NOT IN (?,?,?) " +
		"AND `col_like` LIKE ? " +
		"AND `col_not_like` NOT LIKE ? " +
		"AND `col_like_prefix` LIKE ? " +
		"AND `col_not_like_prefix` NOT LIKE ? " +
		"AND `col_like_suffix` LIKE ? " +
		"AND `col_not_like_suffix` NOT LIKE ? " +
		"AND `col_null` IS NULL " +
		"AND `col_not_null` IS NOT NULL " +
		"AND (`comp_col_eq`=? " +
		"OR `comp_col_neq`!=? " +
		"OR `comp_col_gt`>? " +
		"OR `comp_col_gte`>=? " +
		"OR `comp_col_lt`<? " +
		"OR `comp_col_lte`<=? " +
		"OR `comp_col_between` BETWEEN ? AND ? " +
		"OR `comp_col_not_between` NOT BETWEEN ? AND ? " +
		"OR `comp_col_in` IN (?,?,?) " +
		"OR `comp_col_not_in` NOT IN (?,?,?) " +
		"OR `comp_col_like` LIKE ? " +
		"OR `comp_col_not_like` NOT LIKE ? " +
		"OR `comp_col_like_prefix` LIKE ? " +
		"OR `comp_col_not_like_prefix` NOT LIKE ? " +
		"OR `comp_col_like_suffix` LIKE ? " +
		"OR `comp_col_not_like_suffix` NOT LIKE ? " +
		"OR `comp_col_null` IS NULL " +
		"OR `comp_col_not_null` IS NOT NULL) " +
		"GROUP BY `col_group_by_1`, `col_group_by_2` " +
		"ORDER BY `col_order_by` ASC, `col_order_by_desc` DESC LIMIT 666 OFFSET 999 FOR UPDATE"
	wantParams = []interface{}{"val_eq", "val_neq", "val_gt", "val_gte", "val_lt", "val_lte", "val_between1", "val_between2",
		"val_not_between1", "val_not_between2", "val_in1", "val_in2", "val_in3", "val_not_in1", "val_not_in2", "val_not_in3",
		"%val_like%", "%val_not_like%", "val_like_prefix%", "val_not_like_prefix%", "%val_like_suffix", "%val_not_like_suffix",
		"comp_val_eq", "comp_val_neq", "comp_val_gt", "comp_val_gte", "comp_val_lt", "comp_val_lte", "comp_val_between1",
		"comp_val_between2", "comp_val_not_between1", "comp_val_not_between2", "comp_val_in1", "comp_val_in2",
		"comp_val_in3", "comp_val_not_in1", "comp_val_not_in2", "comp_val_not_in3", "%comp_val_like%", "%comp_val_not_like%",
		"comp_val_like_prefix%", "comp_val_not_like_prefix%", "%comp_val_like_suffix", "%comp_val_not_like_suffix"}

	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)
}

func TestGenerator_SelectByStrct(t *testing.T) {
	m := NewCompExpr()
	m.Put("comp_col_eq", EQ("comp_val_eq"))
	m.Put("comp_col_neq", NEQ("comp_val_neq"))
	m.Put("comp_col_gt", GT("comp_val_gt"))
	m.Put("comp_col_gte", GTE("comp_val_gte"))
	m.Put("comp_col_lt", LT("comp_val_lt"))
	m.Put("comp_col_lte", LTE("comp_val_lte"))
	m.Put("comp_col_between", Between("comp_val_between1", "comp_val_between2"))
	m.Put("comp_col_not_between", NBetween("comp_val_not_between1", "comp_val_not_between2"))
	m.Put("comp_col_in", In([]interface{}{"comp_val_in1", "comp_val_in2", "comp_val_in3"}))
	m.Put("comp_col_not_in", NIn([]interface{}{"comp_val_not_in1", "comp_val_not_in2", "comp_val_not_in3"}))
	m.Put("comp_col_like", Like("comp_val_like"))
	m.Put("comp_col_not_like", NLike("comp_val_not_like"))
	m.Put("comp_col_like_prefix", LikePrefix("comp_val_like_prefix"))
	m.Put("comp_col_not_like_prefix", NLikePrefix("comp_val_not_like_prefix"))
	m.Put("comp_col_like_suffix", LikeSuffix("comp_val_like_suffix"))
	m.Put("comp_col_not_like_suffix", NLikeSuffix("comp_val_not_like_suffix"))
	m.Put("comp_col_null", Null())
	m.Put("comp_col_not_null", NNull())

	ops := []Option{
		WithOr("col_eq", EQ("val_eq")),
		WithOr("col_neq", NEQ("val_neq")),
		WithOr("col_gt", GT("val_gt")),
		WithOr("col_gte", GTE("val_gte")),
		WithOr("col_lt", LT("val_lt")),
		WithOr("col_lte", LTE("val_lte")),
		WithOr("col_between", Between("val_between1", "val_between2")),
		WithOr("col_not_between", NBetween("val_not_between1", "val_not_between2")),
		WithOr("col_in", In([]interface{}{"val_in1", "val_in2", "val_in3"})),
		WithOr("col_not_in", NIn([]interface{}{"val_not_in1", "val_not_in2", "val_not_in3"})),
		WithOr("col_like", Like("val_like")),
		WithOr("col_not_like", NLike("val_not_like")),
		WithOr("col_like_prefix", LikePrefix("val_like_prefix")),
		WithOr("col_not_like_prefix", NLikePrefix("val_not_like_prefix")),
		WithOr("col_like_suffix", LikeSuffix("val_like_suffix")),
		WithOr("col_not_like_suffix", NLikeSuffix("val_not_like_suffix")),
		WithOr("col_null", Null()),
		WithOr("col_not_null", NNull()),
		WithOrExprs(m),
		WithOrderBy("col_order_by"),
		WithOrderByDESC("col_order_by_desc"),
		WithLimit(666),
		WithOffset(999),
		ForceIndex("idx_some_index"),
	}

	wantSQL := "SELECT `col_a`, `col_b` FROM `table_name` FORCE INDEX (`idx_some_index`) " +
		"WHERE `col_eq`=? " +
		"OR `col_neq`!=? " +
		"OR `col_gt`>? " +
		"OR `col_gte`>=? " +
		"OR `col_lt`<? " +
		"OR `col_lte`<=? " +
		"OR `col_between` BETWEEN ? AND ? " +
		"OR `col_not_between` NOT BETWEEN ? AND ? " +
		"OR `col_in` IN (?,?,?) " +
		"OR `col_not_in` NOT IN (?,?,?) " +
		"OR `col_like` LIKE ? " +
		"OR `col_not_like` NOT LIKE ? " +
		"OR `col_like_prefix` LIKE ? " +
		"OR `col_not_like_prefix` NOT LIKE ? " +
		"OR `col_like_suffix` LIKE ? " +
		"OR `col_not_like_suffix` NOT LIKE ? " +
		"OR `col_null` IS NULL " +
		"OR `col_not_null` IS NOT NULL " +
		"OR (`comp_col_eq`=? " +
		"AND `comp_col_neq`!=? " +
		"AND `comp_col_gt`>? " +
		"AND `comp_col_gte`>=? " +
		"AND `comp_col_lt`<? " +
		"AND `comp_col_lte`<=? " +
		"AND `comp_col_between` BETWEEN ? AND ? " +
		"AND `comp_col_not_between` NOT BETWEEN ? AND ? " +
		"AND `comp_col_in` IN (?,?,?) " +
		"AND `comp_col_not_in` NOT IN (?,?,?) " +
		"AND `comp_col_like` LIKE ? " +
		"AND `comp_col_not_like` NOT LIKE ? " +
		"AND `comp_col_like_prefix` LIKE ? " +
		"AND `comp_col_not_like_prefix` NOT LIKE ? " +
		"AND `comp_col_like_suffix` LIKE ? " +
		"AND `comp_col_not_like_suffix` NOT LIKE ? " +
		"AND `comp_col_null` IS NULL " +
		"AND `comp_col_not_null` IS NOT NULL) " +
		"ORDER BY `col_order_by` ASC, `col_order_by_desc` DESC LIMIT 666 OFFSET 999"
	wantParams := []interface{}{"val_eq", "val_neq", "val_gt", "val_gte", "val_lt", "val_lte", "val_between1", "val_between2",
		"val_not_between1", "val_not_between2", "val_in1", "val_in2", "val_in3", "val_not_in1", "val_not_in2", "val_not_in3",
		"%val_like%", "%val_not_like%", "val_like_prefix%", "val_not_like_prefix%", "%val_like_suffix", "%val_not_like_suffix",
		"comp_val_eq", "comp_val_neq", "comp_val_gt", "comp_val_gte", "comp_val_lt", "comp_val_lte", "comp_val_between1",
		"comp_val_between2", "comp_val_not_between1", "comp_val_not_between2", "comp_val_in1", "comp_val_in2",
		"comp_val_in3", "comp_val_not_in1", "comp_val_not_in2", "comp_val_not_in3", "%comp_val_like%", "%comp_val_not_like%",
		"comp_val_like_prefix%", "comp_val_not_like_prefix%", "%comp_val_like_suffix", "%comp_val_not_like_suffix"}

	type structure struct {
		ColA string `db:"col_a"`
		ColB int    `db:"col_b"`
	}

	g := NewGenerator("table_name", ops...)
	_, _, err := g.SelectByStruct(nil)
	assertError(t, err, errors.New("target can not be empty"))

	gotSQL, gotParams, err := g.SelectByStruct(structure{})
	assertError(t, err, nil)
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)

	gotSQL, gotParams, err = g.SelectByStruct(&structure{})
	assertError(t, err, nil)
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)

	var s *structure
	gotSQL, gotParams, err = g.SelectByStruct(s)
	assertError(t, err, nil)
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)
}

func TestGenerator_Update(t *testing.T) {
	opts := []Option{
		WithAnd("col_eq", EQ("val_eq")),
		WithAnd("col_gt", GT("val_gt")),
		WithLimit(1),
	}

	g := NewGenerator("table_name", opts...)
	gotSQL, gotParams := g.Update(nil)
	assertSQL(t, gotSQL, "")
	assertParams(t, gotParams, nil)

	assExpr := NewAssExpr()
	gotSQL, gotParams = g.Update(assExpr)
	assertSQL(t, gotSQL, "")
	assertParams(t, gotParams, nil)

	assExpr.Put("col_1", "val_1")
	assExpr.Put("col_2", "val_2")
	assExpr.Put("col_3", "val_3")
	gotSQL, gotParams = g.Update(assExpr)

	wantSQL := "UPDATE `table_name` SET `col_1`=?, `col_2`=?, `col_3`=? WHERE `col_eq`=? AND `col_gt`>? LIMIT 1"
	wantParams := []interface{}{"val_1", "val_2", "val_3", "val_eq", "val_gt"}
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)
}

func TestGenerator_Delete(t *testing.T) {
	opts := []Option{
		WithAnd("col_eq", EQ("val_eq")),
		WithAnd("col_gt", GT("val_gt")),
		WithLimit(1),
	}

	g := NewGenerator("table_name", opts...)
	gotSQL, gotParams := g.Delete()
	wantSQL := "DELETE FROM `table_name` WHERE `col_eq`=? AND `col_gt`>? LIMIT 1"
	wantParams := []interface{}{"val_eq", "val_gt"}
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)
}

func TestGenerator_Insert(t *testing.T) {
	// INSERT INTO
	g := NewGenerator("table_name")
	columns := []string{"col_1", "col_2", "col_3"}
	records := [][]interface{}{
		{"col_1_1", "col_2_1", "col_3_1"},
		{"col_1_2", "col_2_2", "col_3_2"},
		{"col_1_3", "col_2_3", "col_3_3"},
	}
	gotSQL, gotParams := g.Insert(columns, records...)
	wantSQL := "INSERT INTO `table_name` (`col_1`, `col_2`, `col_3`) VALUES (?,?,?), (?,?,?), (?,?,?)"
	wantParams := []interface{}{"col_1_1", "col_2_1", "col_3_1", "col_1_2", "col_2_2", "col_3_2", "col_1_3", "col_2_3", "col_3_3"}
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)

	// INSERT INTO ON DUPLICATE KEY UPDATE
	assExpr := NewAssExpr()
	assExpr.Put("col_1", "col_1_1")
	assExpr.Put("col_2", "col_1_2")
	opts := []Option{
		OnDuplicateKeyUpdate(assExpr),
	}
	g = NewGenerator("table_name", opts...)
	gotSQL, gotParams = g.Insert(columns, records...)
	wantSQL = "INSERT INTO `table_name` (`col_1`, `col_2`, `col_3`) VALUES (?,?,?), (?,?,?), (?,?,?) " +
		"ON DUPLICATE KEY UPDATE `col_1`=?, `col_2`=?"
	wantParams = []interface{}{
		"col_1_1", "col_2_1", "col_3_1", "col_1_2", "col_2_2", "col_3_2", "col_1_3", "col_2_3", "col_3_3", "col_1_1", "col_1_2"}
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)

	// INSERT INTO WHERE NOT EXIST
	exprCom := NewCompExpr()
	exprCom.Put("col_1", GTE("val_gte"))
	exprCom.Put("col_2", EQ("val_eq"))
	opts = []Option{
		WithNExists("table_name1", exprCom),
	}
	g = NewGenerator("table_name", opts...)
	gotSQL, gotParams = g.Insert(columns, records...)
	wantSQL = "INSERT INTO `table_name` (`col_1`, `col_2`, `col_3`) SELECT ?,?,? " +
		"FROM dual WHERE NOT EXISTS (SELECT * FROM `table_name1` WHERE `col_1`>=? AND `col_2`=?)"
	wantParams = []interface{}{"col_1_1", "col_2_1", "col_3_1", "val_gte", "val_eq"}
	assertSQL(t, gotSQL, wantSQL)
	assertParams(t, gotParams, wantParams)
}

func assertSQL(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got sql dose not meet the expected\nexpected: %s\n  actual: %s", want, got)
	}
}

func assertParams(t *testing.T, got, want []interface{}) {
	if len(got) != len(want) {
		t.Errorf("got params dose not meet the expected\nexpected length: %d\n  actual length: %d", len(want), len(got))
	}

	for i, w := range want {
		if got[i] != w {
			t.Errorf("got params dose not meet the expected\nexpected: %s\n  actual: %s", want, got)
		}
	}
}

func assertError(t *testing.T, got, want error) {
	gotErr := ""
	if got != nil {
		gotErr = got.Error()
	}

	wantErr := ""
	if want != nil {
		wantErr = want.Error()
	}

	if gotErr != wantErr {
		t.Errorf("got error dose not meet the expected\nexpected: %v\n  antual: %v", want, got)
	}
}
