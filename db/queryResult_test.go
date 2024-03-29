package db

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/bigdot123456/golib4db/err1"
)

func TestQueryResultToByte(t *testing.T) {
	r := &res{
		columns: []string{"string",
			"[]byte", "bool", "float64", "float32",
			"int64", "int32"},
		datalength: 2,
		err:        err1.NewError(901, "错误内容"),
		data: [][]interface{}{
			{
				"字符串1",
				[]byte("bytes"),
				false,
				10.23,
				1.2,
				64,
				1,
			},
			{
				"字符串2",
				[]byte("bytes2"),
				true,
				5.1,
				8.3,
				6,
				19,
			},
		},
	}
	b := QueryResultToBytes(r)
	fmt.Println(hex.EncodeToString(b))
	r2 := BytesToQueryResult(b)
	fmt.Println(r2.Columns())
	fmt.Println(r2.Length())
}
