package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

const (
	sqlDefaultSep = ","
)

type Tuple2IntInt struct {
	Ele1 int64 `db:"_1"`
	Ele2 int64 `db:"_2"`
}

func Placeholders(count int) string {
	return appendDuplicateString("?", sqlDefaultSep, count)
}

func appendDuplicateString(character, separator string, count int) string {
	if count <= 0 {
		return ""
	}
	var b bytes.Buffer
	for i := 0; i < count; i++ {
		if i > 0 {
			b.WriteString(separator)
		}
		b.WriteString(character)
	}
	return b.String()
}

func BuildSqlArgs(args ...interface{}) ([]interface{}, error) {
	newArgs := make([]interface{}, 0)
	addEleFun := func(ele interface{}) {
		newArgs = append(newArgs, ele)
		return
	}
	for _, arg := range args {
		switch v := arg.(type) {
		case string, int, int32, int64, bool, *string, *int, *int32, *int64, time.Time:
			addEleFun(v)
		case []string:
			for _, e := range v {
				addEleFun(e)
			}
		case []int:
			for _, e := range v {
				addEleFun(e)
			}
		case []int32:
			for _, e := range v {
				addEleFun(e)
			}
		case []int64:
			for _, e := range v {
				addEleFun(e)
			}
		case []*string:
			for _, e := range v {
				addEleFun(e)
			}
		default:
			return nil, errors.New("类型不匹配")
		}
	}
	return newArgs, nil
}

func MultiInsertValues(columnCount int, count int) string {
	if columnCount <= 0 || count <= 0 {
		return ""
	}
	one := fmt.Sprintf("(%s)", appendDuplicateString("?", sqlDefaultSep, columnCount))
	return appendDuplicateString(one, sqlDefaultSep, count)
}
