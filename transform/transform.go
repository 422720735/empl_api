package transform

import (
	"errors"
	"fmt"
	"strconv"
)

// 接口转string
func InterToString(inter interface{}) (string, error) {
	if inter == nil {
		return "", errors.New("参数错误")
	}
	switch v := inter.(type) {
	case []byte:
		return string(v), nil
	case int:
		str := strconv.Itoa(v)
		return str, nil
	case string:
		return v, nil
	case float64:
		str := strconv.FormatFloat(v, 'f', -1, 64)
		return str, nil
	default:
		return "", fmt.Errorf("type %T", v)
	}
}
