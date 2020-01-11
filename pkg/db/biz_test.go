package db

import (
	"fmt"
	"testing"
)

func TestDollarToQuestionMark(t *testing.T) {
	sql := `select * from users where id = $1, name = $2, server = $3     , stop = $1, tom = $2`

	adaptSQL, args, err := dollarToQuestionMark(sql, "aaaaa", "b", "ccc")
	if err != nil {
		fmt.Printf("adaptSQL = '%s'", adaptSQL)
		fmt.Printf("args = '%v'", args)
		t.Error(err)
	}
	trueSQL := `select * from users where id = ?, name = ?, server = ?     , stop = ?, tom = ?`
	if adaptSQL != trueSQL {
		fmt.Println("adaptSQL error")
		fmt.Println("True is :", trueSQL)
		fmt.Println("Real is :", adaptSQL)
		t.Error()
	}
	if len(args) != 5 {
		t.Error("args len not valid")
	}
}
